package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
	"net"
	"strconv"
	"summer/server/cmd/analyze/config"
	"summer/server/cmd/task/initialize"
	"summer/server/cmd/task/pkg"
	"summer/server/shared/consts"
	task "summer/server/shared/kitex_gen/task/taskservice"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	IP, Port := initialize.InitFlag()
	rc := initialize.InitRedis()
	ac := initialize.InitAnalyze()

	r, info := initialize.InitRegistry(Port)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	impl := &TaskServiceImpl{
		RedisManger:   pkg.NewRedisManager(rc),
		AnalyzeManger: pkg.NewAnalyzeManger(ac),
	}

	svr := task.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
