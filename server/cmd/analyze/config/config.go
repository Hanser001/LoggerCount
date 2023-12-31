package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key" json:"secret_access_key"`
	Bucket          string `mapstructure:"bucket" json:"bucket"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Host        string        `mapstructure:"host" json:"host"`
	MinioInfo   MinioConfig   `mapstructure:"minio" json:"minio"`
	OtelInfo    OtelConfig    `mapstructure:"otel" json:"otel"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
