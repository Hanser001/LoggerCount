package config

import "summer/server/shared/kitex_gen/user/userservice"

var (
	GlobalServerConfig ServerConfig
	GlobalConsulConfig ConsulConfig

	UserClient userservice.Client
)
