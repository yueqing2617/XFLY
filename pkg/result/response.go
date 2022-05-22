package e

import "github.com/gin-gonic/gin"

// 简化response
type Resp struct {
	Ctx *gin.Context
}

// Response 方法
func (r *Resp) Response(code int, data any) {
	r.Ctx.JSON(200, gin.H{
		"code":    code,
		"data":    data,
		"message": GetMsg(code),
	})
	return
}

func (r Resp) Fatail(code int, message string) {
	if message == "" {
		message = GetMsg(code)
	} else {
		message = GetMsg(code) + ": " + message
	}
	r.Ctx.JSON(500, gin.H{
		"code":    code,
		"data":    nil,
		"message": "message",
	})
	return
}
