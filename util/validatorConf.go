package util

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

//设置Validator提示中文
func SetValidatorZH() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		//注册自定义函数
		//_=v.RegisterValidation("bookabledate", bookableDate)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
		//根据提供的标记注册翻译
		//v.RegisterTranslation("bookabledate", trans, func(ut ut.Translator) error {
		//	return ut.Add("bookabledate", "{0}不能早于当前时间或{1}格式错误!", true)
		//}, func(ut ut.Translator, fe validator.FieldError) string {
		//	t, _ := ut.T("bookabledate", fe.Field(), fe.Field())
		//	return t
		//})

	}
}

func TranslateToSting(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	var build strings.Builder
	if ok {
		errorsTranslations := errs.Translate(trans)
		for _, val := range errorsTranslations {
			build.WriteString(val)
			build.WriteString("|")
		}
	}
	return build.String()
}
