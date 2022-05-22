package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yueqing2617/XFLY/internal/routes"
)

// 初始化路由
func Init() *gin.Engine {
	fmt.Println("正在启动服务")
	// load routes
	r := routes.LoadRoutes()
	// laad static files
	r.StaticFS("./storage", http.Dir("storage"))
	fmt.Println("服务启动成功")
	// 返回路由对象
	return r
}
