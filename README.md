## 技术选型

- HTTP框架Hertz
- RPC框架Kitex
- 服务中心与配置中心选用Consul
- 关系型数据库MySQL
- 非关系型数据库Redis
- 分布式文件系统选用Minio
- 消息队列RabbitMQ
- 使用 Jaeger 与 Prometheus 进行链路追踪以及监控
- 访问控制模型Casbin
- 使用sentinel进行熔断限流

## 架构设计

### 调用关系

![image-20230618092518846](https://typora-1314425967.cos.ap-nanjing.myqcloud.com/typora/image-20230618092518846.png)

### 数据库设计

MySQL：

- 维护User表
- Casbin使用Gorm Adaptor接入了Mysql，维护casbin_rules表

Redis:

- 为上传文件的用户维护一个sorted set，用户id作为key，文件名作为filed，上传时间作为score
- 为运行过数据分析任务的用户维护一个hash，用户id作为key，任务id作为字段名，value为0或1

## 项目代码介绍

### 项目代码结构

#### 整体结构

```
├── docker-compose.yaml
├── otel-collector-config.yaml
├── go.mod
├── go.sum
├── server
│   ├── cmd
│   │   ├── api
│   │   ├── analyze
│   │   ├── task
│   │   ├── user
│   │   └── file
│   ├── idl
│   │   ├── api.thrift
│   │   ├── errno.thrift
│   │   ├── analyze.thrift
│   │   ├── file.thrift
│   │   ├── user.thrift
│   │   └── task.thrift
│   └── shared
│       ├── consts
│       ├── errno
│       ├── kitex_gen
│       ├── middleware
│       └── tools
```

#### 微服务内部结构

以user服务为例

```
├── config
│   └── config.go
│   └── global.go
├── config.yaml
├── dao
│   ├── user.go
├── handler.go
├── initialize
│   ├── config.go
│   ├── registry.go
│   ├── flag.go
│   ├── logger.go
│   ├── db.go
├── kitex_info.yaml
├── main.go
├── model
│   └── user.go
└── pkg
    ├── md5.go
```

#### 日志拓展

```go
// InitLogger to init logrus
func InitLogger() {
	// Customizable output directory.
	logFilePath := consts.KlogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := kitexlogrus.NewLogger()
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	if runtime.GOOS == "linux" {
		logger.SetOutput(lumberjackLogger)
		logger.SetLevel(klog.LevelWarn)
	} else {
		logger.SetLevel(klog.LevelDebug)
	}

	klog.SetLogger(logger)
}

```

所有日志拓展均选用logrus，包括gorm，kitex，hertz。

#### 对象存储Minio

```go
//InitMinio to init minio
func InitMinio() *minio.Client {
	config := config2.GlobalServerConfig.MinioInfo
	// Initialize minio client object.
	mc, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		klog.Fatalf("create minio client err: %s", err.Error())
	}
	exists, err := mc.BucketExists(context.Background(), config.Bucket)
	if err != nil {
		klog.Fatal(err)
	}
	if !exists {
		err = mc.MakeBucket(context.Background(), config.Bucket, minio.MakeBucketOptions{Region: "cn-north-1"})
		if err != nil {
			klog.Fatalf("make bucket err: %s", err.Error())
		}
	}

	return mc
}

```

当用户需要下载文件，后台会返回一个为期七天的对应资源的URL

```go
func (s *FileServiceImpl) DownloadFile(ctx context.Context, req *file.DownloadRequest) (resp *file.DownloadResponse, err error) {
	resp = new(file.DownloadResponse)

	bucketName := config.GlobalServerConfig.MinioInfo.Bucket
	url, err := s.Dao.Download(ctx, bucketName, req.ObjectName)
	if err != nil {
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		resp.Url = ""
		return resp, nil
	}

	resp.StatusCode = int32(errno.Success.ErrCode)
	resp.StatusMsg = errno.Success.ErrMsg
	resp.Url = url.String()

	return resp, nil
}
```

#### 数据处理

1.用户发起请求后，Task Srv会调用Analyze Srv。

```go
func (s *TaskServiceImpl) NewTask_(ctx context.Context, req *task.NewTaskRequest_) (resp *task.NewTaskResponse_, err error) {
    ...
	//ignore logic that creates new task

	taskId := time.Now().Unix()
	s.RedisManger.SetTaskRecord(ctx, int(req.UserId), int(taskId))

	go s.AnalyzeManger.Analyze(ctx, "example1", "example2", "example3", req.UserId)

    ...
}
```

2.从Minio取得文件并把数据丢到我们缝合得到的用于WordCount的mapreduce框架进行处理。

```go
unc StartWordCount(files []string, filed string) ([]*KV, error) {
	if len(files) == 0 {
		return nil, errors.New("no data")
	}
	wc := NewMaster(files)
	middleData, err := wc.Map(filed)
	if err != nil {
		return nil, err
	}

	groups := wc.Generalize(middleData)

	result := wc.Reduce(groups)

	return result, nil
}
```

3.得到处理结果后会自动上传到minio，Redis也会更新数据。

```go
func (m *RedisManger) SetTaskRecord(ctx context.Context, userId int, taskId int) error {
	// TODO: add error check
	m.client.HSet(ctx, strconv.Itoa(userId),
		strconv.Itoa(taskId),
		consts.TaskWorking,
	)
	return nil
}
```

#### 中间件

##### Gzip

```go
gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".jpg", ".log"})),
```

使用 Hertz提供的Gzip 中间件资源进行压缩，并自定义不进行压缩的资源格式。

Gzip能优化web应用性能。

##### Casbin

在Web应用，可以将路由作为obj，请求方式Method作为act，登录用户的角色作为sub，在每一次请求时，把这3个参数传递给e.Enforce， 就可以实现对Web页面和请求接口的权限控制管理，非常方便。

这里使用RBAC with Pattern模型，我将group作为sub替代了登录用户角色，并为用户写入对应grouping policy。

```
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && regexMatch(r.act, p.act)
```

在创建用户时会给用户添加Grouping Policy

```go
e := casbin.GetEnforcer()
	_, err = e.AddGroupingPolicy(usr.ID, "user")
```

通过hertz的casbin拓展，用 `NewCasbinMiddlewareFromEnforcer` 初始化casbin，再用 `RequiresRoles` 中间件方法返回app.Handlerfunc并注册到对应路由。

```go
func MyCasbinAuth(role string) app.HandlerFunc {
	cas := InitCasbin()
	handlerfunc := cas.RequiresRoles(role)
	return handlerfunc
}
```

##### Limiter

使用 hertz的Limiter 中间件对项目进行限流

```go
limiter.AdaptiveLimit(limiter.WithCPUThreshold(900)),
```

#### 服务治理

##### Consul

Consul同时承担配置中心和服务中心。

```go
// InitRegistry to init consul
func InitRegistry() (registry.Registry, *registry.Info)
```

```go
// InitConfig to init consul config server
func InitConfig()
```

关于Conusl KV配置，[详见](https://github.com/Hanser001/LoggerCount/blob/master/docs/Consul%20KV%20Config.md)。

##### Opentelemetry

race 使用 Jaeger，Metrics 使用 Prometheus，Logs 使用 Logrus。

```go
p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
```

#### 安全

##### JWT

JWT密钥保存在配置中心，实现脱敏。

在对应路由使用JWTAuth中间件进行鉴权。

```go
func _uploadMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middwares.JWTAuth(config.GlobalServerConfig.JWTInfo.SigningKey),
		casbin.MyCasbinAuth(consts.UserRole),
	}
}
```

##### MD5

用户注册与登录输入密码均要通过MD5加密。

```
// Md5Crypt uses MD5 encryption algorithm to add salt encryption.
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

```

与JWT密钥相同，MD5的salt也保存在配置中心。

## 坑点与收获

### 收获

- 狠狠地学习了Docker，也是第一次使用微服务框架
- 掌握了一些流行中间件的基本使用
- 提高了阅读文档能力，对于Casbin，Minio API使用和Hertz中间件的学习文档都有很大帮助

### 坑点

- gorm的logrus拓展跟kitex，hertz的logrus拓展：如果拉取最新版本 `go.opentelemetry.io/otel`会缺失依赖，需要回退版本
- 对分布式计算框架还是点都不懂，项目实现的最简单的用于WordCount的mapreduce框架还是缝合来的

## 参考

[Freecar](https://github.com/CyanAsterisk/FreeCar)

[TikGok](https://github.com/CyanAsterisk/TikGok)