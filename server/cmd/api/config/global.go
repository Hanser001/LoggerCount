package config

import (
	"summer/server/shared/kitex_gen/analyze/analyzeservice"
	"summer/server/shared/kitex_gen/file/fileservice"
	"summer/server/shared/kitex_gen/task/taskservice"
	"summer/server/shared/kitex_gen/user/userservice"
)

var (
	GlobalConsulConfig ConsulConfig
	GlobalServerConfig ServerConfig

	GlobalUserClient    userservice.Client
	GlobalFileClient    fileservice.Client
	GlobalTaskClient    taskservice.Client
	GlobalAnalyzeClient analyzeservice.Client
)
