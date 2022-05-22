package model

import (
	"github.com/yueqing2617/XFLY/internal/model/system_model"
	"xorm.io/xorm"
)

func LoadModel(Db *xorm.Engine) {
	// 加载system_model 目录下的model
	Db.Sync2(new(system_model.Roles), new(system_model.Admin), new(system_model.Menu))
}
