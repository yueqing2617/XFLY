package material_model

import (
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MaterialLarghezza
// Description: Materiale di larghezza
// Table: material_larghezza
// Fields: id, larghezza
// Author: zhoudongming
// Time: 2022/5/02 15:15 下午

type MaterialLarghezza struct {
	Id        int64  `xorm:"pk autoincr comment('自增ID')" json:"id"`
	Larghezza string `xorm:"varchar(255) notnull comment('材料的门幅')" json:"larghezza"`
}

// TableName 表名
func (MaterialLarghezza) TableName() string {
	return "material_larghezza"
}

// TableObject 表对象
func (this MaterialLarghezza) TableObject() *xorm.Session {
	return db.XormEngine.Table(this.TableName())
}

// GetAll 获取所有
func (this MaterialLarghezza) GetAll() ([]MaterialLarghezza, error) {
	var list []MaterialLarghezza
	err := this.TableObject().Find(&list)
	return list, err
}

// GetById 根据id获取
func (this MaterialLarghezza) GetById(id int64) (MaterialLarghezza, error) {
	var item MaterialLarghezza
	has, err := this.TableObject().Where("id = ?", id).Get(&item)
	if err != nil {
		return item, err
	}
	if !has {
		return item, nil
	}
	return item, nil
}

// Add 添加
func (this MaterialLarghezza) Add() (int64, error) {
	id, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}
