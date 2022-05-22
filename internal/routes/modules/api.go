package modules

import (
	"github.com/gin-gonic/gin"
	e "github.com/yueqing2617/XFLY/pkg/result"
)

func LoadApiRoutes(route *gin.Engine) {
	app := route.Group("/api")
	{
		app.GET("/", func(c *gin.Context) {
			h := e.Resp{Ctx: c}
			h.Response(e.Success, "")
		})
	}
}
