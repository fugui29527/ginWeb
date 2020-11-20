package controller

import (
	"encoding/base64"
	"fmt"
	"ginWeb/dto"
	"ginWeb/models"
	"ginWeb/service"
	"ginWeb/util"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"time"
)

type UserController struct {
}

//注冊
func (this *UserController) Register(c *gin.Context) {
	userRegister := new(dto.UserRegister)
	c.BindJSON(userRegister)
	userservice := new(service.UserService)
	u := models.CreateUserInfo(userRegister.Username, userRegister.PassWord, userRegister.Sex)
	_, err := userservice.Insert(u)
	if err != nil {
		util.Faile(c, "注冊失敗")
		return
	}
	util.Success(c, u)
}

var configCharacter = base64Captcha.ConfigCharacter{
	Height: 60,
	Width:  100,
	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	Mode:               base64Captcha.CaptchaModeNumber,
	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         4,
}

//生成验证码
func (this *UserController) GenerateCode(c *gin.Context) {
	//var store = base64Captcha.DefaultMemStore
	var captchaImageChar = base64Captcha.EngineCharCreate(configCharacter)
	verifyValue := captchaImageChar.VerifyValue
	fmt.Println("====验证码:" + verifyValue)
	//生成imageBase64 字符串
	base64 := base64.StdEncoding.EncodeToString(captchaImageChar.BinaryEncoding())

	//保存到缓存
	valKey := "validate:code:" + "12345678"
	set := util.SetNx(valKey, verifyValue, 60*time.Second)
	fmt.Println("redis 保存结果:", set)
	util.Success(c, gin.H{
		"code":        captchaImageChar.VerifyValue,
		"base64Image": base64,
		"id":          "12345678",
	})
}

//登陆
func (this *UserController) Login(c *gin.Context) {

}
