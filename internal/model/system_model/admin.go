package system_model

import (
	"errors"
	"log"

	"github.com/yueqing2617/XFLY/pkg/utils"
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: Admin
// Description: Admin model
// Fields: id, username, password, email, telephone, avatar, nickname, department, description, status, role_id, created_time, updated_time,  gender
// Author: zhoudongming
// Time: 2022/05/01

type Admin struct {
	Id          int64  `xorm:"pk autoincr comment('自增ID')" json:"id"`
	Username    string `xorm:"varchar(50) notnull comment('用户名')" json:"username"`
	Password    string `xorm:"varchar(50) notnull comment('密码')" json:"password"`
	Email       string `xorm:"varchar(50) comment('邮箱')" json:"email"`
	Telephone   string `xorm:"varchar(50) notnull comment('电话')" json:"telephone"`
	Avatar      string `xorm:"varchar(50) comment('头像')" json:"avatar"`
	Nickname    string `xorm:"varchar(50) comment('昵称')" json:"nickname"`
	Gender      int    `xorm:"int(1) notnull comment('性别') default(1)"` // 1:男 2:女
	Department  string `xorm:"varchar(50) comment('部门')" json:"department"`
	Description string `xorm:"varchar(50) comment('描述')" json:"description"`
	Status      int    `xorm:"int(11) notnull comment('状态') default(1)" json:"status"` // 1:正常, 2:禁用, 3:删除
	RoleId      int64  `xorm:"int(11) notnull comment('角色ID')" json:"role_id"`
	CreatedTime int64  `xorm:"created comment('创建时间')" json:"created_time"`
	UpdatedTime int64  `xorm:"updated comment('更新时间')" json:"updated_time"`
}

// TableName 表名
func (Admin) TableName() string {
	return "admin"
}

// TableObject 表对象
func (Admin) TableObject() *xorm.Session {
	return db.XormEngine.Table("admin")
}

// Add 添加
func (a *Admin) Add() (int64, error) {
	_, err := a.TableObject().Insert(a)
	if err != nil {
		log.Fatalln("添加管理员失败:", err)
		return 0, err
	}
	return a.Id, nil
}

// Update 更新给定的字段
func (a *Admin) Update(fields ...string) error {
	m := make(map[string]interface{})
	data := utils.StructToMap(m)
	for _, field := range fields {
		m[field] = data[field]
	}
	_, err := a.TableObject().Where("id = ?", a.Id).Update(m)
	if err != nil {
		log.Fatalln("更新管理员失败:", err)
		return err
	}
	return nil
}

// Delete 删除(软删除)
func (a *Admin) Delete() error {
	if a.Status == 3 {
		return errors.New("该管理员已被删除")
	}
	if a.Id == 1 {
		return errors.New("不能删除超级管理员")
	}
	_, err := a.TableObject().Where("id = ?", a.Id).Cols("status").Update(a)
	if err != nil {
		return errors.New("删除管理员失败")
	}
	return nil
}

// GetById 获取给定ID的管理员
func (a *Admin) GetById(id int64) (*Admin, error) {
	_, err := a.TableObject().Where("id = ?", id).Get(a)
	if err != nil {
		log.Fatalln("获取管理员失败:", err)
		return nil, err
	}
	return a, nil
}

// Login 登录
func (a *Admin) Login(username, password string) (*Admin, error) {
	// 判断用户名或手机号是否存在
	_, err := a.TableObject().Where("username = ?", username).Or("telephone = ?", username).Get(a)
	if err != nil {
		log.Fatalln("登录失败:", err)
		return nil, err
	}
	// 判断密码是否正确
	if !utils.PasswordCompare(password, a.Password) {
		return nil, errors.New("密码错误")
	}
	return a, nil
}

// ExistUsername 判断管理员用户名是否存在
func (a *Admin) ExistUsername(username string) (bool, error) {
	return a.TableObject().Where("username = ?", username).Exist()
}

// ExistTelephone 判断管理员手机号是否存在
func (a *Admin) ExistTelephone(telephone string) (bool, error) {
	return a.TableObject().Where("telephone = ?", telephone).Exist()
}

// Register 注册
func (a *Admin) Register() error {
	// 密码加密
	a.Password = utils.PasswordEncrypt(a.Password)
	// 添加管理员
	_, err := a.Add()
	if err != nil {
		log.Fatalln("注册管理员失败:", err)
		return err
	}
	return nil
}
