package material_model

import (
	"errors"
	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MaterialSample
// Description: 布料样品
// DatabaseName: material_sample
// Fileds: id,name,code,type,status,remark, element, larghezza, gram_wight, specification, description, color, ImageList,created_time, updated_time
// Author: zhoudongming
// Time: 2022/05/01 15:14:00

type MaterialSample struct {
	Id            int64  `xorm:"pk autoincr comment('自增ID')" json:"id"`
	Name          string `xorm:"comment('名称')" json:"name"`
	Code          string `xorm:"comment('编码')" json:"code"`
	Type          string `xorm:"comment('类型')" json:"type"`
	Status        int    `xorm:"not null default 0 comment('状态: 1-启用, 3-已删除')" json:"status"`
	Remark        string `xorm:"text comment('备注')" json:"remark"`
	Element       string `xorm:"text comment('成分')" json:"element"`
	Larghezza     string `xorm:"text comment('门幅')" json:"larghezza"`
	GramWight     string `xorm:"text comment('克重')" json:"gram_wight"`
	Specification string `xorm:"text comment('规格')" json:"specification"`
	Description   string `xorm:"text comment('描述')" json:"description"`
	Color         string `xorm:"text comment('颜色')" json:"color"`
	ImageList     string `xorm:"text comment('图片')" json:"image_list"`
	CreatedTime   int64  `xorm:"created comment('创建时间')" json:"created_time"`
	UpdatedTime   int64  `xorm:"updated comment('更新时间')" json:"updated_time"`
}

// TableName 表名
func (MaterialSample) TableName() string {
	return "material_sample"
}

// TableObject 设置引擎
func (this MaterialSample) TableObject() *xorm.Session {
	return db.XormEngine.Table(this.TableName())
}

// Add 新增
func (this MaterialSample) Add() (int64, error) {
	id, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新给定的字段
func (this MaterialSample) Update(fields ...string) error {
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

// Delete 删除(软删除)
func (this MaterialSample) Delete() error {
	if this.Status == 3 {
		return errors.New("已删除")
	}
	this.Status = 3
	_, err := this.TableObject().Where("id = ?", this.Id).Update(this)
	if err != nil {
		return err
	}
	return nil
}

// GetMaterialSampleList 根据条件获取列表,支持分页。
// 返回 total, items, err
func (this MaterialSample) GetMaterialSampleList(page int, pageSize int, where string, args ...interface{}) (int64, []MaterialSample, error) {
	var list []MaterialSample
	var total int64
	var err error
	if pageSize <= 0 {
		pageSize = 10
	}
	if page <= 0 {
		page = 1
	}
	if pageSize == -1 {
		total, err = this.TableObject().Where(where, args...).Count()
		if err != nil {
			return 0, nil, err
		}
		err = this.TableObject().Where(where, args...).Find(&list)
	} else {
		total, err = this.TableObject().Where(where, args...).Count()
		if err != nil {
			return 0, nil, err
		}
		err = this.TableObject().Where(where, args...).Limit(pageSize, (page-1)*pageSize).Find(&list)
	}
	if err != nil {
		return 0, nil, err
	}
	return total, list, nil
}
