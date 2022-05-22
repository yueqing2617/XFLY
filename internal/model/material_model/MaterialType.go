package material_model

import (
	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MaterialType
// Description: MaterialType Model
// Fields: id, name, description, parentId
// Author: zhoudongming
// Time: 2022/5/02

type MaterialType struct {
	Id          int64          `xorm:"pk autoincr comment('自增ID')" json:"id"`
	Name        string         `xorm:"varchar(50) notnull comment('布料分类名')" json:"name"`
	Description string         `xorm:"varchar(50) comment('描述')" json:"description"`
	ParentId    int64          `xorm:"int(11) comment('父级分类ID') default(0)" json:"parent_id"`
	Children    []MaterialType `xorm:"-" json:"children"`
}

// TableName 表名
func (this MaterialType) TableName() string {
	return "material_type"
}

// TableObject 表对象
func (this MaterialType) TableObject() *xorm.Session {
	return db.XormEngine.Table("material_type")
}

// GetMaterialTypeList 获取物料类型树型列表
func (this MaterialType) GetMaterialTypeList() ([]MaterialType, error) {
	var material_types []MaterialType
	if err := this.TableObject().Find(&material_types); err != nil {
		return nil, err
	}
	return material_types, nil
}

// GetMaterialTypeTreeList 获取物料类型树型列表
func (this MaterialType) GetMaterialTypeTreeList() (data []MaterialType, err error) {
	var material_types []MaterialType
	if err = this.TableObject().Find(&material_types); err != nil {
		return nil, err
	}
	data = listToTree(material_types, 0)
	return data, nil
}

// listToTree 将菜单列表转换为树形结构
func listToTree(data []MaterialType, parentId int64) []MaterialType {
	var tree []MaterialType
	for _, item := range data {
		if item.ParentId == parentId {
			item.Children = listToTree(data, item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}

// Add 添加
func (this MaterialType) Add() (int64, error) {
	id, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新给定的字段
func (this MaterialType) Update(fields ...string) error {
	m := make(map[string]interface{})
	data := utils.StructToMap(m)
	for _, field := range fields {
		m[field] = data[field]
	}
	_, err := this.TableObject().Where("id=?", this.Id).Update(m)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (this MaterialType) Delete() error {
	_, err := this.TableObject().Where("id=?", this.Id).Delete(this)
	if err != nil {
		return err
	}
	return nil
}
