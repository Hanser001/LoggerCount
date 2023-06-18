package casbin

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	hertzcasbin "github.com/hertz-contrib/casbin"
	c "summer/server/cmd/user/config"
	"summer/server/shared/consts"
)

// GetEnforcer get enforcer instance
func GetEnforcer() *casbin.Enforcer {
	config := c.GlobalServerConfig.MysqlInfo
	dsn := GetDsn(config)
	//adapter, err := gormadapter.NewAdapter("mysql", dsn)

	adapter, err := gormadapter.NewAdapter("mysql", dsn)
	if err != nil {
		klog.Fatal("casbin err")
	}

	enforcer, err := casbin.NewEnforcer("./model.conf", adapter)
	if err != nil {
		klog.Fatal("casbin err")
	}

	return enforcer
}

func GetDsn(c c.MysqlConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		c.Name,
		c.Password,
		c.Host,
		c.Port,
	)
}

func InitCasbin() *hertzcasbin.Middleware {
	e := GetEnforcer()

	casbinMiddleware, err := hertzcasbin.NewCasbinMiddlewareFromEnforcer(e, MyLookupHandler)
	if err != nil {
		klog.Fatal("casbin middleware err")
	}

	return casbinMiddleware
}

func MyCasbinAuth(role string) app.HandlerFunc {
	cas := InitCasbin()
	handlerfunc := cas.RequiresRoles(role)
	return handlerfunc
}

func MyLookupHandler(ctx context.Context, c *app.RequestContext) string {
	uid := c.GetString(consts.AccountID)
	//belong to uid,get role from gorm adaptor
	e := GetEnforcer()
	ok := e.HasNamedPolicy("g", uid, "user")

	if !ok {
		return ""
	}
	return "user"
}
