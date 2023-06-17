package config

import "github.com/minio/minio-go/v7"

var (
	GlobalServerConfig ServerConfig
	GlobalConsulConfig ConsulConfig

	MinioClient minio.Client
)
