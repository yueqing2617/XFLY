package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yueqing2617/XFLY/service/conf"
	"xorm.io/xorm"
)

// InitXormDB 初始化xorm
var XormEngine *xorm.Engine

func InitXormDB(conf *conf.DatabaseConfig) *xorm.Engine {
	log.Println("读取数据库配置")
	// dsn
	dsn := ""
	if conf.Type == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Suffix)
	} else if conf.Type == "sqlite3" {
		dsn = fmt.Sprintf("%s/%s", conf.Host, conf.Database)
	}
	// 初始化xorm
	var err error
	XormEngine, err = xorm.NewEngine(conf.Type, dsn)
	if err != nil {
		log.Fatalln("数据库连接失败: ", err)
	}
	if conf.Debug {
		XormEngine.ShowSQL(true)
	}
	log.Println("读取数据库配置成功")
	return XormEngine
}

// 获取xorm
func GetXorm() *xorm.Engine {
	return XormEngine
}
