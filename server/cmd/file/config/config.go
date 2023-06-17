package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key" json:"secret_access_key"`
	Bucket          string `mapstructure:"bucket" json:"bucket"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type RabbitMqConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Exchange string `mapstructure:"exchange" json:"exchange"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Name         string         `mapstructure:"name" json:"name"`
	Host         string         `mapstructure:"host" json:"host"`
	MinioInfo    MinioConfig    `mapstructure:"minio" json:"minio"`
	OtelInfo     OtelConfig     `mapstructure:"otel" json:"otel"`
	RabbitMqInfo RabbitMqConfig `mapstructure:"rabbitmq" json:"rabbitmq"`
	RedisInfo    RedisConfig    `mapstructure:"redis" json:"redis"`
	UserSrvInfo  UserSrvConfig  `mapstructure:"user_srv" json:"user_srv"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
