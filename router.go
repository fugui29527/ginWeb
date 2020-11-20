package main

import (
	"ginWeb/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(app *gin.Engine) {
	//用户相关
	userGroup := app.Group("/user")
	user := new(controller.UserController)
	userGroup.POST("/register", user.Register)
	userGroup.GET("/generateCode", user.GenerateCode)
	userGroup.POST("/login", user.Login)
}
