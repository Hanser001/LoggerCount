package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name           string           `mapstructure:"name" json:"name"`
	Host           string           `mapstructure:"host" json:"host"`
	RedisInfo      RedisConfig      `mapstructure:"redis" json:"redis"`
	OtelInfo       OtelConfig       `mapstructure:"otel" json:"otel"`
	AnalyzeStvInfo AnalyzeSrvConfig `mapstructure:"analyze_srv" json:"analyze_srv"`
	FileSrvInfo    FileSrvConfig    `mapstructure:"file_srv" json:"file_srv"`
}

type AnalyzeSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type FileSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
