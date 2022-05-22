package common_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yueqing2617/XFLY/internal/model/system_model"
	"github.com/yueqing2617/XFLY/pkg/captcha"
	"github.com/yueqing2617/XFLY/pkg/jwt"
	e "github.com/yueqing2617/XFLY/pkg/result"
	"github.com/yueqing2617/XFLY/pkg/utils"
)

type Common struct{}

func NewCommon() *Common {
	return &Common{}
}

// @Summary Login
// @Description 登录
// @Tags common
// @Accept  json
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param code body string true "验证码"
// @Param appid body string true "appid"
// @Success 200 string string "{"code":200,"data":{},"msg":"ok"}"
// @Router /common/login [post]
func (c *Common) Login(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	type form struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Code     string `json:"code"`
		Appid    string `json:"appid"`
	}
	var params form
	if err := ctx.ShouldBindJSON(&params); err != nil {
		h.Fatail(e.BadRequest, err.Error())
	}
	// 验证表单
	if params.Username == "" {
		h.Fatail(e.BadRequest, "用户名不能为空")
	}
	if params.Password == "" {
		h.Fatail(e.BadRequest, "密码不能为空")
	}
	if params.Code == "" {
		h.Fatail(e.BadRequest, "验证码不能为空")
	}
	if params.Appid == "" {
		h.Fatail(e.BadRequest, "非法请求")
	}
	// 验证码
	if captcha.VerifyCaptcha(params.Appid, params.Code) != true {
		h.Fatail(e.BadRequest, "验证码错误")
	}
	// 登录
	admin, err := new(system_model.Admin).Login(params.Username, params.Password)
	if err != nil {
		h.Fatail(e.BadRequest, err.Error())
	}
	// 验证成功，生成token
	j := jwt.NewJwt()
	token, err := j.MakeToken(admin.Id, admin.RoleId)
	if err != nil {
		h.Fatail(e.InternalServerError, err.Error())
	}
	h.Response(e.Success, gin.H{"token": token})
}

// @Summary Register
// @Description 注册
// @Tags common
// @Accept  json
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param telphone body string true "手机号"
// @Param nickname body string true "昵称"
// @Param ConfirmPassword body string true "确认密码"
// @Param email body string false "邮箱"
// @Param avatar body string false "头像"
// @Param department body string false "部门"
// @Param gender body int false "性别"
// @Param code body string true "验证码"
// @Param appid body string true "appid"
// @Success 200 string string "{"code":200,"data":{},"msg":"ok"}"
// @Router /common/register [post]
func (c *Common) Register(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	type form struct {
		system_model.Admin
		ConfirmPassword string `json:"confirmPassword"`
		Code            string `json:"code"`
		Appid           string `json:"appid"`
	}
	var params form
	if err := ctx.ShouldBindJSON(&params); err != nil {
		h.Fatail(e.BadRequest, err.Error())
	}
	// 验证表单
	if captcha.VerifyCaptcha(params.Appid, params.Code) != true {
		h.Fatail(e.BadRequest, "验证码错误")
	}
	if params.Username == "" {
		h.Fatail(e.BadRequest, "用户名不能为空")
	}
	if params.Password == "" {
		h.Fatail(e.BadRequest, "密码不能为空")
	}
	if params.ConfirmPassword == "" {
		h.Fatail(e.BadRequest, "确认密码不能为空")
	}
	if params.Telephone == "" {
		h.Fatail(e.BadRequest, "手机号不能为空")
	}
	if params.Nickname == "" {
		h.Fatail(e.BadRequest, "昵称不能为空")
	}
	if params.Code == "" {
		h.Fatail(e.BadRequest, "验证码不能为空")
	}
	if params.Appid == "" {
		h.Fatail(e.BadRequest, "非法请求")
	}
	if params.Password != params.ConfirmPassword {
		h.Fatail(e.BadRequest, "密码不一致")
	}
	// 验证用户名是否存在
	isp, err := params.ExistUsername(params.Username)
	if err != nil || isp {
		h.Fatail(e.BadRequest, "用户名已存在")
	}
	// 验证手机号是否存在
	isp, err = params.ExistTelephone(params.Telephone)
	if err != nil || isp {
		h.Fatail(e.BadRequest, "手机号已存在")
	}
	if err := params.Register(); err != nil {
		h.Fatail(e.InternalServerError, err.Error())
	}
	h.Response(e.Success, nil)
}

// @Summary GetCaptcha
// @Description 获取验证码
// @Tags common
// @Accept  json
// @Produce  json
// @Success 200 string string "{"code":200,"data":{},"msg":"ok"}"
// @Router /common/captcha [get]
func (c *Common) GetCaptcha(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	// generate captcha code
	appid, code, err := captcha.MakeCaptcha()
	if err != nil {
		h.Fatail(e.InternalServerError, err.Error())
	}
	h.Response(e.Success, gin.H{"rand": appid, "code": code})
}

// @Summary Test
// @Description 测试
// @Tags common
// @Accept  json
// @Produce  json
// @Success 200 string string "{"code":200,"data":{},"msg":"ok"}"
// @Router /common/test [get]
func (c *Common) Test(ctx *gin.Context) {
	h := e.Resp{Ctx: ctx}
	file, err := utils.ReadNetFile("http://dh.inzj.cn/test.txt")
	if err != nil {
		h.Fatail(e.InternalServerError, err.Error())
	}
	// 将 file 解析为 txt 文件
	txt := utils.ParseTxt(file)
	h.Response(e.Success, gin.H{"text": txt})
}
