package material_model

import (
	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MaterialSpecification
// Description: MaterialSpecification
// DatabaseName: material_specifications
// Fields: id, specification
// Author: zhoudongming
// Time: 2022/05/02 16:23:04

type MaterialSpecification struct {
	Id            int64  `xorm:"pk autoincr comment('自增ID')" json:"id"`
	Specification string `xorm:"comment('规格')" json:"specification"`
}

// TableName 表名
func (MaterialSpecification) TableName() string {
	return "material_specifications"
}

// TableObject 设置引擎
func (this MaterialSpecification) TableObject() *xorm.Session {
	return db.XormEngine.Table(this.TableName())
}

// Add 新增
func (this MaterialSpecification) Add() (int64, error) {
	id, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新给定的字段
func (this MaterialSpecification) Update(fields ...string) error {
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
func (this MaterialSpecification) Delete() error {
	_, err := this.TableObject().Where("id = ?", this.Id).Delete(this)
	return err
}

// GetById 根据id获取
func (this MaterialSpecification) GetById() (MaterialSpecification, error) {
	var data MaterialSpecification
	has, err := this.TableObject().Where("id = ?", this.Id).Get(&data)
	if err != nil {
		return data, err
	}
	if !has {
		return data, nil
	}
	return data, nil
}

// GetAll 获取所有
func (this MaterialSpecification) GetAll() ([]MaterialSpecification, error) {
	var datas []MaterialSpecification
	err := this.TableObject().Find(&datas)
	if err != nil {
		return datas, err
	}
	return datas, nil
}
