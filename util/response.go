package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 200
	FAILE   = -1
)

//返回结果信息
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/**
  成功返回
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
		Code: FAILE,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, result)
}
