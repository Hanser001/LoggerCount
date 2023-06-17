package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/kitex/pkg/klog"
	c "summer/server/cmd/user/config"
)

// GetEnforcer get enforcer instance
func GetEnforcer() *casbin.Enforcer {
	//config := c.GlobalServerConfig.MysqlInfo
	//dsn := GetDsn(config)
	//adapter, err := gormadapter.NewAdapter("mysql", dsn)

	adapter, err := gormadapter.NewAdapter("mysql", "root:Wjj20040311!@tcp(39.101.68.42:3306)/")
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
