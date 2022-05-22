package e

var MsgFlags = map[int]string{
	Success:          "获取成功",
	Error:            "获取失败",
	BadRequest:       "请求参数错误",
	Unauthorized:     "未授权",
	Forbidden:        "禁止访问",
	NotFound:         "资源不存在",
	MethodNotAllowed: "请求方法不允许",
	RequestTimeout:   "请求超时",
}

// GetMsg get message
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
