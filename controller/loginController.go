package controller

import (
	"ginWeb/dto"
	"ginWeb/service"
	"ginWeb/util"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

// @ 获取图形验证码
// @Summary 获取图形验证码
// @Tags login
// @Version v1.0
// @Description 获取图形验证码
// @Accept  json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /generateCode [get]
func (this *LoginController) GenerateCode(c *gin.Context) {
	//var store = base64Captcha.DefaultMemStore
	idKey, base64Pic := util.GenerateCaptcha()
	util.Success(c, gin.H{
		"base64Image": base64Pic,
		"idKey":       idKey,
	})
}

// @ 登录接口
// @Summary 登录接口
// @Tags login
// @Version v1.0
// @Description 登录接口
// @Accept  application/json
// @Produce json
// @Param data body   dto.LoginDto      true     "请求参数"
// @Success 200 {object} util.Response 	"ok"
// @Failure 999 object string 失败
// @Router /login [Post]
func (this *LoginController) Login(c *gin.Context) {

	loginDto := new(dto.LoginDto)
	if err := c.ShouldBindJSON(loginDto); err != nil {
		errMsg := util.TranslateToSting(err)
		util.FaileCode(c, util.PARAM_FAILE, errMsg)
		return
	}
	//验证图形嘛是否正确
	if result := util.VerifyCaptcha(loginDto.IdKey, loginDto.VailCode); !result {
		util.FaileCode(c, util.CAPTCHA_FAILE, "图形验证码错误!")
		return
	}
	service := new(service.LoginService)
	token, err := service.DoLogin(loginDto)
	if err != nil {
		util.Faile(c, err.Error())
		return
	}
	util.Success(c, token)
}
