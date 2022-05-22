package system_model

import (
	"github.com/yueqing2617/XFLY/service/db"
	"xorm.io/xorm"
)

// Model: Menu
// Description: 权限菜单
// Fields: id, name, url, icon, parent_id, sort, status, menu_type, code, operator, created_time, updated_time
type Menu struct {
	Id       int64  `xorm:"pk autoincr comment('菜单ID')" json:"id"`
	Name     string `xorm:"varchar(50) notnull comment('菜单名称')" json:"label"`
	Path     string `xorm:"text comment('菜单Path')" json:"path"`
	Method   string `xorm:"varchar(50) comment('菜单访问方法')" json:"method"`
	ParentId int64  `xorm:"int(11) comment('父级菜单ID')" json:"parent_id"`
	Children []Menu `xorm:"-" json:"children"`
}

// TableName 表名
func (Menu) TableName() string {
	return "admin_permission_menu"
}

// TableObject 表对象
func (Menu) TableObject() *xorm.Session {
	return db.XormEngine.Table("admin_permission_menu")
}

// GetMenuList 获取菜单列表
func (m Menu) GetMenuList() ([]Menu, error) {
	var menus []Menu
	err := m.TableObject().Find(&menus)
	if err != nil {
		return nil, err
	}
	data := listToTree(menus, 0)
	return data, nil
}

// listToTree 将菜单列表转换为树形结构
// parentId: 父级ID
func listToTree(menus []Menu, parentId int64) []Menu {
	var tree []Menu
	for _, menu := range menus {
		if menu.ParentId == parentId {
			menu.Children = listToTree(menus, menu.Id)
			tree = append(tree, menu)
		}
	}
	return tree
}
