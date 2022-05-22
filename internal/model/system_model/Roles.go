package system_model

import (
	"errors"
	"log"

	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: Roles
// Description: Roles
// Fields: id, name, description, status, alias, created_time, updated_time

type Roles struct {
	Id            int64  `xorm:"autoincr pk comment('主键id')" json:"id"`
	Name          string `xorm:"varchar(50) notnull comment('角色名称')" validata:"required" json:"name"`
	Description   string `xorm:"varchar(255) comment('角色描述')" json:"description"`
	Status        int    `xorm:"tinyint(1) notnull default(1) comment('状态：1-正常，2-禁用，3-删除')" json:"status"`
	Alias         string `xorm:"varchar(50) notnull comment('角色别名')" json:"alias"`
	Authorization string `xorm:"text comment('拥有权限')" json:"authorization"`
	CreatedTime   int64  `xorm:"created comment('创建时间')" json:"created_time"`
	UpdatedTime   int64  `xorm:"updated comment('更新时间')" json:"updated_time"`
}

// TableName 表名
func (Roles) TableName() string {
	return "admin_roles"
}

// TableObject 表对象
func (Roles) TableObject() *xorm.Session {
	return db.XormEngine.Table("admin_roles")
}

// add 添加
func (r *Roles) Add() (int64, error) {
	id, err := r.TableObject().Insert(r)
	if err != nil {
		log.Fatalln("添加角色失败：", err)
		return 0, err
	}
	return id, nil
}

// update 更新给定的字段
func (r *Roles) Update(fields ...string) error {
	m := make(map[string]interface{})
	data := utils.StructToMap(m)
	for _, field := range fields {
		m[field] = data[field]
	}
	_, err := r.TableObject().Where("id = ?", r.Id).Update(m)
	if err != nil {
		log.Fatalln("更新管理员失败:", err)
		return err
	}
	return nil
}

// delete 删除(软删除)
func (r *Roles) Delete() error {
	if r.Status == 3 {
		return errors.New("角色已删除")
	}
	if r.Id == 1 {
		return errors.New("超级管理员不能删除")
	}
	_, err := r.TableObject().Where("id = ?", r.Id).Cols("status").Update(r)
	if err != nil {
		log.Fatalln("删除角色失败:", err)
		return err
	}
	return nil
}

// getById 获取给定id的角色
func (r *Roles) GetById(id int64) (*Roles, error) {
	_, err := r.TableObject().Where("id = ?", id).Get(r)
	if err != nil {
		log.Fatalln("获取角色失败:", err)
		return nil, err
	}
	return r, nil
}

// getByAlias 获取给定别名的角色
func (r *Roles) GetByAlias(alias string) (*Roles, error) {
	_, err := r.TableObject().Where("alias = ?", alias).Get(r)
	if err != nil {
		log.Fatalln("获取角色失败:", err)
		return nil, err
	}
	return r, nil
}

// ExistByAlias 是否存在给定别名的角色
func (r *Roles) ExistByAlias(alias string) bool {
	count, err := r.TableObject().Where("alias = ?", alias).Count()
	if err != nil {
		log.Fatalln("获取角色失败:", err)
		return false
	}
	return count > 0
}

// GetAll 获取所有角色,并根据传入的page,limit进行分页
func (r *Roles) GetAll(page, limit int) (data []*Roles, count int64, err error) {
	var roles []*Roles
	count, _ = r.TableObject().Where("status != 3").Count()
	err = r.TableObject().Limit(limit, (page-1)*limit).Find(&roles)
	if err != nil {
		log.Fatalln("获取角色失败:", err)
		return nil, 0, err
	}
	return roles, count, nil
}
