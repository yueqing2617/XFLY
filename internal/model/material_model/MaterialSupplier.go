package material_model

import (
	"errors"
	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MaterialSupplier
// Description: MaterialSupplier
// Database table: material_supplier
// Fields: id, name, buyer, phone, address, status, created_time, updated_time
// Author: zhoudongming
// Time: 2022/05/02 15:00

type MaterialSupplier struct {
	Id          int64  `xorm:"pk autoincr comment('自增ID')" json:"id"`
	Name        string `xorm:"not null comment('供应商名称')" json:"name"`
	Buyer       string `xorm:"not null comment('联系人')" json:"buyer"`
	Phone       string `xorm:"comment('联系电话')" json:"phone"`
	Address     string `xorm:"comment('地址')" json:"address"`
	Status      int    `xorm:"not null default 0 comment('状态: 1-启用, 3-禁用')" json:"status"`
	CreatedTime int64  `xorm:"created comment('创建时间')" json:"created_time"`
	UpdatedTime int64  `xorm:"updated comment('更新时间')" json:"updated_time"`
}

// TableName 表名
func (this MaterialSupplier) TableName() string {
	return "material_supplier"
}

// TableObject 设置引擎
func (this MaterialSupplier) TableObject() *xorm.Session {
	return db.XormEngine.Table(this.TableName())
}

// Add 添加
func (this MaterialSupplier) Add() (int64, error) {
	id, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新给定的字段
func (this MaterialSupplier) Update(fields ...string) error {
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

// Delete 删除(软删除)，将status置为3
func (this MaterialSupplier) Delete() error {
	if this.Status == 3 {
		return errors.New("已经删除")
	}
	this.Status = 3
	_, err := this.TableObject().Where("id = ?", this.Id).Update(this)
	if err != nil {
		return err
	}
	return nil
}

// GetById 根据id获取
func (this MaterialSupplier) GetById() (MaterialSupplier, error) {
	var materialSupplier MaterialSupplier
	has, err := this.TableObject().Where("id = ?", this.Id).Get(&materialSupplier)
	if err != nil {
		return materialSupplier, err
	}
	if !has {
		return materialSupplier, errors.New("没有找到")
	}
	return materialSupplier, nil
}

// GetByName 根据name获取
func (this MaterialSupplier) GetByName() (MaterialSupplier, error) {
	var materialSupplier MaterialSupplier
	has, err := this.TableObject().Where("name = ?", this.Name).Get(&materialSupplier)
	if err != nil {
		return materialSupplier, err
	}
	if !has {
		return materialSupplier, errors.New("没有找到")
	}
	return materialSupplier, nil
}

// GetAll 获取所有
func (this MaterialSupplier) GetAll() ([]MaterialSupplier, error) {
	var materialSuppliers []MaterialSupplier
	err := this.TableObject().Find(&materialSuppliers)
	if err != nil {
		return materialSuppliers, err
	}
	return materialSuppliers, nil
}

// GetMaterialSupplierrList 根据条件获取列表,支持分页。
// 返回 total, items, err
func (this MaterialSupplier) GetMaterialSupplierrList(page int, pageSize int, where string, args ...interface{}) (int64, []MaterialSupplier, error) {
	var materialSuppliers []MaterialSupplier
	session := this.TableObject().Where(where, args...)
	total, err := session.Count()
	if err != nil {
		return 0, materialSuppliers, err
	}
	if total == 0 {
		return 0, materialSuppliers, nil
	}
	if pageSize > 0 {
		session.Limit(pageSize, (page-1)*pageSize)
	}
	err = session.Desc("id").Find(&materialSuppliers)
	if err != nil {
		return 0, materialSuppliers, err
	}
	return total, materialSuppliers, nil
}
