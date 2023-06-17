package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"net"
	"strconv"
	"summer/server/cmd/file/config"
	"summer/server/cmd/file/dao"
	"summer/server/cmd/file/initialize"
	"summer/server/cmd/file/pkg"
	"summer/server/shared/consts"
	file "summer/server/shared/kitex_gen/file/fileservice"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	mc := initialize.InitMinio()
	IP, Port := initialize.InitFlag()
	rdb := initialize.InitRedis()

	r, info := initialize.InitRegistry(Port)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	impl := &FileServiceImpl{
		Dao:         dao.NewFileManger(mc),
		RedisManger: pkg.NewRedisManager(rdb),
	}

	svr := file.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
