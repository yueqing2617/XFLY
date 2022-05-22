package material_model

import (
	"errors"
	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: MaterialPurchaseOrder
// Description: MaterialPurchaseOrder
// Database table: material_purchase_orders
// Fields: id, material_id, supplier_id, quantity, price, unit, total, status, created_time, updated_time
// Associations: material, supplier
// BelongsTo: material, supplier
// Author: zhoudongming
// Time: 2022/05/02 16:23 下午

type MaterialPurchaseOrder struct {
	Id          int64   `xorm:"pk autoincr comment('自增ID')" json:"id"`
	MaterialId  int64   `xorm:"not null comment('物料ID')" json:"material_id"`
	SupplierId  int64   `xorm:"not null comment('供应商ID')" json:"supplier_id"`
	Quantity    int64   `xorm:"comment('数量')" json:"quantity"`
	Price       float64 `xorm:"comment('单价')" json:"price"`
	Unit        string  `xorm:"comment('单位')" json:"unit"`
	Total       float64 `xorm:"comment('总价')" json:"total"`
	Status      int64   `xorm:"not null comment('状态: 1-正常, 2-已完成, 3-已取消')" json:"status"`
	CreatedTime int64   `xorm:"created comment('创建时间')" json:"created_time"`
	UpdatedTime int64   `xorm:"updated comment('更新时间')" json:"updated_time"`
}

// TableName 表名
func (this MaterialPurchaseOrder) TableName() string {
	return "material_purchase_orders"
}

// TableObject 设置引擎
func (this MaterialPurchaseOrder) TableObject() *xorm.Session {
	return db.XormEngine.Table(this.TableName())
}

// Add 新增
func (this MaterialPurchaseOrder) Add() (int64, error) {
	id, err := this.TableObject().Insert(this)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新给定的字段
func (this MaterialPurchaseOrder) Update(fields ...string) error {
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

// Delete 删除,软删除，将status改为3
func (this MaterialPurchaseOrder) Delete() error {
	if this.Status == 3 {
		return errors.New("该订单已取消")
	}
	if this.Status == 2 {
		return errors.New("该订单已完成, 不能取消")
	}
	this.Status = 3
	_, err := this.TableObject().Where("id=?", this.Id).Update(this)
	if err != nil {
		return err
	}
	return nil
}

// GetById 根据id获取
func (this MaterialPurchaseOrder) GetById(id int64) (MaterialPurchaseOrder, error) {
	var order MaterialPurchaseOrder
	has, err := this.TableObject().Where("id=?", id).Get(&order)
	if err != nil {
		return order, err
	}
	if !has {
		return order, errors.New("没有找到该订单")
	}
	return order, nil
}

// GetMaterialPurchaseOrderList 根据条件获取列表,支持分页。
// 返回 total, items, err
func (this MaterialPurchaseOrder) GetMaterialPurchaseOrderList(page int, pageSize int, where string, args ...interface{}) (int64, []MaterialPurchaseOrder, error) {
	var list []MaterialPurchaseOrder
	total, err := this.TableObject().Where(where, args...).Count()
	if err != nil {
		return 0, nil, err
	}
	err = this.TableObject().Where(where, args...).Limit(pageSize, (page-1)*pageSize).Desc("id").Find(&list)
	if err != nil {
		return 0, nil, err
	}
	return total, list, nil
}

// GetAllMaterialPurchaseOrderList 获取全部列表
func (this MaterialPurchaseOrder) GetAllMaterialPurchaseOrderList() ([]MaterialPurchaseOrder, error) {
	var list []MaterialPurchaseOrder
	err := this.TableObject().Desc("id").Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
