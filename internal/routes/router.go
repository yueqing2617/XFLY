package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yueqing2617/XFLY/internal/routes/modules"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Get the module name in the modules folder
func LoadRoutes() *gin.Engine {
	// Load the routes
	r := gin.New()
	// include the options
	Include(modules.LoadAdminRoutes, modules.LoadApiRoutes)
	for _, opt := range options {
		opt(r)
	}

	return r
}
