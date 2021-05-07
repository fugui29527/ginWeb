package util

import "github.com/mojocn/base64Captcha"

var configCharacter = base64Captcha.ConfigCharacter{
	Height: 40,
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

//生成base64图片
func GenerateCaptcha() (idKey string, base64Pic string) {

	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKey, capC := base64Captcha.GenerateCaptcha("", configCharacter)
	//以base64编码
	base64Pic = base64Captcha.CaptchaWriteToBase64Encoding(capC)
	return
}

//验证base64图片
func VerifyCaptcha(idKey string, validateCode string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(idKey, validateCode)
	return verifyResult
}
