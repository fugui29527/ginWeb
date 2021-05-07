package controller

import (
	"ginWeb/util"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

//生成验证码
func (this *LoginController) GenerateCode(c *gin.Context) {
	//var store = base64Captcha.DefaultMemStore
	idKey, base64Pic := util.GenerateCaptcha()
	util.Success(c, gin.H{
		"base64Image": base64Pic,
		"idKey":       idKey,
	})
}
