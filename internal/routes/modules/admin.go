package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/yueqing2617/XFLY/internal/controller/common_controller"
	"github.com/yueqing2617/XFLY/internal/controller/system_controller"
	e "github.com/yueqing2617/XFLY/pkg/result"
)

func LoadAdminRoutes(route *gin.Engine) {
	app := route.Group("/admin")
	{
		app.GET("/", func(ctx *gin.Context) {
			h := e.Resp{Ctx: ctx}
			h.Response(e.Success, "")
		})
		// common routes
		common := app.Group("/common")
		{
			commons := common_controller.NewCommon()
			// login
			common.POST("/login", commons.Login)
			// Register
			common.POST("/register", commons.Register)
			// getcaptcha
			common.GET("/captcha", commons.GetCaptcha)
			// test
			common.GET("/test", commons.Test)
		}
		// roles routes
		roles := app.Group("/roles")
		{
			rolesController := system_controller.NewRoles()
			// add role
			roles.POST("/add", rolesController.Add)
			// edit role
			roles.POST("/edit", rolesController.Edit)
			// delete role
			roles.POST("/delete", rolesController.Del)
			// get role list
			roles.GET("/list", rolesController.List)
			// get role info
			roles.GET("/:id", rolesController.Info)
		}
		// premissions routes
		premissions := app.Group("/premissions")
		{
			premissionsController := system_controller.NewPremission()
			//// add premissions
			//premissions.POST("/add", premissionsController.Add)
			//// edit premissions
			//premissions.POST("/edit", premissionsController.Edit)
			//// delete premissions
			//premissions.POST("/delete", premissionsController.Del)
			// get premissions list
			premissions.GET("/list", premissionsController.List)
		}
	}
}
