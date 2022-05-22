package system_controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yueqing2617/XFLY/internal/model/system_model"
	e "github.com/yueqing2617/XFLY/pkg/result"
	"github.com/yueqing2617/XFLY/pkg/utils"
)

type Roles struct{}

func NewRoles() *Roles {
	return &Roles{}
}

// @Summary 新增角色
// @Description 新增角色
// @Accept json
// @Produce json
// @Param role body system_controller.Role true "role"
// @Success 200 {object} system_controller.Role
// @Failure 400 {object} system_controller.Error
// @Router /admin/roles/add [post]
func (this *Roles) Add(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	var role system_model.Roles
	if err := ctx.ShouldBindJSON(&role); err != nil {
		h.Fatail(e.BadRequest, err.Error())
		return
	}

	if _, err := role.Add(); err != nil {
		h.Fatail(e.InternalServerError, err.Error())
		return
	}

	h.Response(e.Success, nil)
}

// @Summary 删除角色
// @Description 删除角色
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 string "ok"
// @Failure 400 {object} system_controller.Error
// @Router /admin/roles/del/{id} [del]
func (this *Roles) Del(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	strId := ctx.Param("id")
	id, _ := strconv.Atoi(strId)
	if id == 0 {
		h.Fatail(e.BadRequest, "ID不能为空")
		return
	}
	role, err := new(system_model.Roles).GetById(int64(id))
	if err != nil {
		h.Fatail(e.InternalServerError, err.Error())
		return
	}
	if role == nil {
		h.Fatail(e.Error, "角色不存在")
		return
	}
	if err := role.Delete(); err != nil {
		h.Fatail(e.InternalServerError, err.Error())
		return
	}
	h.Response(e.Success, nil)
}

// @Summary 更新角色
// @Description 更新角色
// @Accept json
// @Produce json
// @Param role body system_controller.Role true "role"
// @Success 200 string "ok"
// @Failure 400 string "err"
// @Router /admin/roles/edit [post]
func (this *Roles) Edit(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	var role system_model.Roles
	if err := ctx.ShouldBindJSON(&role); err != nil {
		h.Fatail(e.BadRequest, err.Error())
		return
	}
	if role.Id == 0 {
		h.Fatail(e.BadRequest, "ID不能为空")
		return
	}
	if err := role.Update("name", "description", "status"); err != nil {
		h.Fatail(e.InternalServerError, err.Error())
		return
	}
	h.Response(e.Success, nil)
}

// @Summary 获取角色列表
// @Description 获取角色列表
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} system_controller.Roles
// @Failure 400 string "err"
// @Router /admin/roles/list [get]
func (this *Roles) List(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	roles, count, err := new(system_model.Roles).GetAll(page, limit)
	if err != nil {
		h.Fatail(e.InternalServerError, err.Error())
		return
	}
	h.Response(e.Success, gin.H{
		"count": count,
		"data":  roles,
	})
}

// @Summary 获取角色详情
// @Description 获取角色详情
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} system_controller.Roles
// @Failure 400 string "err"
// @Router /admin/roles/{id} [get]
func (this *Roles) Info(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	idStr := ctx.Param("id")

	id := utils.StringToInt64(idStr)
	if id == 0 {
		h.Fatail(e.BadRequest, "ID不能为空")
		return
	}
	role, err := new(system_model.Roles).GetById(id)
	if err != nil {
		h.Fatail(e.InternalServerError, err.Error())
		return
	}
	if role.Id == 0 {
		h.Fatail(e.Error, "角色不存在")
		return
	}
	h.Response(e.Success, gin.H{
		"data": role,
	})
}
