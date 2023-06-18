package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type RabbitMqConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Exchange string `mapstructure:"exchange" json:"exchange"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ServerConfig struct {
	Name           string       `mapstructure:"name" json:"name"`
	Host           string       `mapstructure:"host" json:"host"`
	Port           int          `mapstructure:"port" json:"port"`
	OtelInfo       OtelConfig   `mapstructure:"otel" json:"otel"`
	JWTInfo        JWTConfig    `mapstructure:"jwt" json:"jwt"`
	UserSrvInfo    RPCSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	FileSrvInfo    RPCSrvConfig `mapstructure:"file_srv" json:"file_srv"`
	TaskSrvInfo    RPCSrvConfig `mapstructure:"task_srv" json:"task_srv"`
	AnalyzeSrvInfo RPCSrvConfig `mapstructure:"analyze_srv" json:"analyze_srv"`
}

type RPCSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
