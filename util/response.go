package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS       int = 200
	PARAM_FAILE   int = 201
	CAPTCHA_FAILE int = 202
	FAILE         int = 999
)

//返回结果信息
type Response struct {
	Code int         `json:"code" example:"状态码"`
	Msg  string      `json:"msg" example:"信息"`
	Data interface{} `json:"data" `
}

/**
b不带参数返回成功
*/
func SuccessNoData(c *gin.Context) {
	result := &Response{
		Code: SUCCESS,
		Data: nil,
	}
	c.JSON(http.StatusOK, result)
}

/**
  成功返回带参数
*/
func Success(c *gin.Context, data interface{}) {
	result := &Response{
		Code: SUCCESS,
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}

/**
失败返回
*/
func Faile(c *gin.Context, msg string) {
	result := &Response{
		Code: FAILE,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, result)
}

/**
帶code返回
*/
func FaileCode(c *gin.Context, code int, msg string) {
	result := &Response{
		Code: code,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, result)
}
