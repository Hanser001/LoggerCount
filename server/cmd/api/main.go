// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"
	"net/http"
	"summer/server/cmd/api/config"
	"summer/server/cmd/api/initialize"
	"summer/server/cmd/api/initialize/rpc"
)

func main() {
	// initialize
	initialize.InitLogger()
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	initialize.InitSentinel()
	//tracer, trcCfg := hertztracing.NewServerTracer()
	corsCfg := initialize.InitCors()
	rpc.Init()

	// create a new server
	h := server.New(
		//tracer,
		server.WithHostPorts(fmt.Sprintf(":%d", config.GlobalServerConfig.Port)),
		server.WithRegistry(r, info),
		server.WithHandleMethodNotAllowed(true),
	)

	h.Use(cors.New(corsCfg))
	//h.Use(hertztracing.ServerMiddleware(trcCfg))
	h.Use(hertzSentinel.SentinelServerMiddleware(
		// abort with status 429 by default
		hertzSentinel.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(http.StatusTooManyRequests, nil)
			ctx.Abort()
		}),
	))

	register(h)
	h.Spin()
}