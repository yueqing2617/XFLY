package material_model

import (
	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MaterialElement
// Description: MaterialElement Xorm Model
// Fields: id, name, description, Material
// Author: zhoudongming
// Time: 2022/5/02 14:00 下午

type MaterialElement struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	Name        string `xorm:"varchar(100) notnull comment('名称')" json:"name"`
	Description string `xorm:"varchar(255) comment('描述')" json:"description"`
	Material    int64  `xorm:"notnull comment('材料')" json:"material"`
}

// TableName 表名
func (MaterialElement) TableName() string {
	return "material_element"
}

// TableObject 表对象
func (this MaterialElement) TableObject() *xorm.Session {
	return db.XormEngine.Table(this.TableName())
}

// GetMaterialElementById 根据id获取材料元素
func GetMaterialElementById(id int64) (MaterialElement, error) {
	var materialElement MaterialElement
	has, err := db.XormEngine.ID(id).Get(&materialElement)
	if err != nil {
		return materialElement, err
	}
	if !has {
		return materialElement, nil
	}
	return materialElement, nil
}

// GetMaterialElementList 根据条件获取材料元素列表
func GetMaterialElementList(page, pageSize int) ([]MaterialElement, error) {
	var materialElements []MaterialElement
	err := db.XormEngine.Limit(pageSize, (page-1)*pageSize).Find(&materialElements)
	if err != nil {
		return materialElements, err
	}
	return materialElements, nil
}

// GetMaterialElementCount 根据条件获取材料元素数量
func GetMaterialElementCount() (int64, error) {
	return db.XormEngine.Count(new(MaterialElement))
}

// Add 添加材料元素,返回添加的材料元素id
func (this *MaterialElement) Add() (int64, error) {
	id, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新给定的字段
func (this *MaterialElement) Update(fields ...string) error {
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

// Delete 删除材料元素
func (this *MaterialElement) Delete() error {
	_, err := this.TableObject().Where("id = ?", this.Id).Delete(new(MaterialElement))
	if err != nil {
		return err
	}
	return nil
}

// GetAllMaterialElements 获取所有材料元素
func (this MaterialElement) GetAllMaterialElements() ([]MaterialElement, error) {
	var materialElements []MaterialElement
	err := this.TableObject().Find(&materialElements)
	if err != nil {
		return materialElements, err
	}
	return materialElements, nil
}
