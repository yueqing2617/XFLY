package system_controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yueqing2617/XFLY/internal/model/system_model"
	e "github.com/yueqing2617/XFLY/pkg/result"
)

type Premission struct{}

func NewPremission() *Premission {
	return &Premission{}
}

// @Summary 权限列表
// @Description 权限列表
// @Tags 权限树
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Router /system/permission/list [get]
func (this *Premission) List(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	log.Println("Premission.List")
	data, err := new(system_model.Menu).GetMenuList()
	if err != nil {
		h.Fatail(e.InternalServerError, err.Error())
		return
	}
	h.Response(e.Success, gin.H{
		"data": data,
	})
}
