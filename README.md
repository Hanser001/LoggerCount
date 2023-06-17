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
- 使用sentinel配合limter进行熔断限流

### 数据库设计

Mysql设计:

User表存储用户信息。

Casbin的policy也通过Gorm Adaptor储存在mysql。

Redis设计:

为每一个开启过数据处理任务的用户维护一个Hash，key为user_id，field为task_id，value为0(已完成)或1(未完成)。

为每个上传过文件的用户维护一个Sorted Set，key为user_id，filed为文件id，score为上传时间时间戳。

上传文件时的objname与field都由时间戳生成，方便查询。

### 中间件使用

Casbin使用:根据RBAD with pattern模型，在用户注册时为用户写入grouping policy，利用hertz-contrib中的casbin中间件拓展进行鉴权。

Opentelementry接入：使用Jaeger和Prometheus 进行链路追踪以及监控。

**readme晚上重写！**

## 参考

[Freecar](https://github.com/CyanAsterisk/FreeCar)

[TikGok](https://github.com/CyanAsterisk/TikGok)