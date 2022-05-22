package material_model

import (
	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MarerialGramWeight
// Description: Material model for gram weight
// Fields: id, gram_weight
// Author: zhoudongming
// Time: 2022/5/22 15:00 下午

type MarerialGramWeight struct {
	Id         int64  `xorm:"pk autoincr comment('自增ID')" json:"id"`
	GramWeight string `xorm:"varchar(255) notnull" json:"gram_weight"`
}

// TableName 表名
func (MarerialGramWeight) TableName() string {
	return "material_gram_weight"
}

// TableObject 表对象
func (this MarerialGramWeight) TableObject() *xorm.Session {
	return db.XormEngine.Table(this.TableName())
}

// GetAll 获取所有
func (this MarerialGramWeight) GetAll() ([]MarerialGramWeight, error) {
	var gramWeight []MarerialGramWeight
	err := this.TableObject().Find(&gramWeight)
	return gramWeight, err
}

// GetById 获取单个
func (this MarerialGramWeight) GetById(id int64) (MarerialGramWeight, error) {
	var gramWeight MarerialGramWeight
	has, err := this.TableObject().Where("id = ?", id).Get(&gramWeight)
	if err != nil {
		return gramWeight, err
	}
	if !has {
		return gramWeight, nil
	}
	return gramWeight, nil
}

// Add 添加
func (this MarerialGramWeight) Add() (int64, error) {
	_, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return this.Id, nil
}

// Update 更新给定字段
func (this MarerialGramWeight) Update(fields ...string) error {
	m := make(map[string]interface{})
	data := utils.StructToMap(m)
	for _, field := range fields {
		m[field] = data[field]
	}
	_, err := this.TableObject().Where("id = ?", this.Id).Update(m)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (this MarerialGramWeight) Delete() error {
	_, err := this.TableObject().Where("id = ?", this.Id).Delete(this)
	if err != nil {
		return err
	}
	return nil
}
