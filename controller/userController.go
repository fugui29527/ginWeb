package controller

import (
	"ginWeb/dto"
	"ginWeb/log"
	"ginWeb/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
}

var logger *zap.Logger

func init() {
	logger = log.Logger
}

//注冊
func (this *UserController) Register(c *gin.Context) {
	userRegister := new(dto.UserRegister)
	c.BindJSON(userRegister)

	util.SuccessNoData(c)
}

//登陆
func (this *UserController) Login(c *gin.Context) {

}
