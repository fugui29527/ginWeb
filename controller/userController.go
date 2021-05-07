package controller

import (
	"ginWeb/dto"
	"ginWeb/log"
	"ginWeb/models"
	"ginWeb/service"
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
	userservice := new(service.UserService)
	u := models.CreateUserInfo(userRegister.Username, userRegister.PassWord, userRegister.Sex)
	_, err := userservice.Insert(u)
	if err != nil {
		util.Faile(c, "注冊失敗")
		return
	}
	util.Success(c, u)
}

//登陆
func (this *UserController) Login(c *gin.Context) {

}
