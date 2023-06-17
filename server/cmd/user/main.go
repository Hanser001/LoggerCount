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
	"summer/server/cmd/user/config"
	"summer/server/cmd/user/dao"
	"summer/server/cmd/user/initialize"
	"summer/server/shared/consts"
	user "summer/server/shared/kitex_gen/user/userservice"
	"summer/server/shared/middwares"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	db := initialize.InitDB()
	IP, Port := initialize.InitFlag()

	r, info := initialize.InitRegistry(Port)
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	impl := &UserServiceImpl{
		Jwt: middwares.NewJWT(config.GlobalServerConfig.JWTInfo.SigningKey),
		Dao: dao.NewUserManger(db),
	}

	svr := user.NewServer(impl,
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
