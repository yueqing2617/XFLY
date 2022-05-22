package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/yueqing2617/XFLY/service/db"
)

var Enforcer *casbin.Enforcer

// InitCasbin 初始化casbin
func InitCasbin() {
	log.Println("正在连接数据库")
	a, err := xormadapter.NewAdapterByEngine(db.GetXorm())
	if err != nil {
		log.Fatalln("连接数据库失败", err)
	}
	e, err := casbin.NewEnforcer("./config/rbac_models.conf", a)
	if err != nil {
		log.Fatalln("连接数据库失败", err)
	}
	Enforcer = e
	log.Println("连接数据库成功")
}

// 获取casbin
func GetCasbin() *casbin.Enforcer {
	return Enforcer
}
