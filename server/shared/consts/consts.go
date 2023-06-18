package consts

const (
	MySqlDSN    = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	RabbitMqURI = "amqp://%s:%s@%s:%d/"

	JWTIssuer        = "nanacho"
	ThirtyDays       = 60 * 60 * 24 * 30
	AuthorizationKey = "token"
	Claims           = "claims"
	AccountID        = "accountID"

	ApiConfigPath     = "./server/cmd/api/config.yaml"
	UserConfigPath    = "./server/cmd/user/config.yaml"
	FileConfigPath    = "./server/cmd/file/config.yaml"
	TaskConfigPath    = "./server/cmd/task/config.yaml"
	AnalyzeConfigPath = "./server/cmd/analyze/config.yaml"

	ConsulCheckInterval                       = "7s"
	ConsulCheckTimeout                        = "5s"
	ConsulCheckDeregisterCriticalServiceAfter = "15s"

	TCP             = "tcp"
	FreePortAddress = "localhost:0"

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	ConsulSnowflakeNode = 1
	UserSnowflakeNode   = 2
	FileSnowflakeNode   = 3
	MinioSnowflakeNode  = 4

	IPFlagName  = "ip"
	IPFlagValue = "0.0.0.0"
	IPFlagUsage = "address"

	PortFlagName  = "port"
	PortFlagUsage = "port"

	RedisProfileClientDB = 0

	UrlExpiredTime = 1800

	NewLocalFilePath = "./tmp/%d.log"

	Lines = 2000

	TaskSuccess = 0
	TaskWorking = 1

	CorsAddress = "http://localhost:3000"

	UserRole = "user"
)
