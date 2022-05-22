package service

import (
	"log"

	"github.com/yueqing2617/XFLY/internal/model"
	"github.com/yueqing2617/XFLY/service/cache"
	"github.com/yueqing2617/XFLY/service/casbin"
	"github.com/yueqing2617/XFLY/service/conf"
	"github.com/yueqing2617/XFLY/service/db"
	"github.com/yueqing2617/XFLY/service/routers"
)

func Run() {
	// 加载配置
	conf.InitConfig()
	// 加载数据库
	sql := db.InitXormDB(conf.GetConfig().DatabaseConfig)
	// 迁移数据库
	model.LoadModel(sql)
	// 加载缓存
	cache.InitCache(conf.GetConfig().CacheConfig)
	// 加载Casbin
	casbin.InitCasbin()
	// 加载路由
	run := routers.Init()
	// 启动服务
	if err := run.Run(":8888"); err != nil {
		log.Fatalln("启动服务失败", err)
	}
}
