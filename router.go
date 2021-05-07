package main

import (
	"ginWeb/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(app *gin.Engine) {
	//登录相关
	login := new(controller.LoginController)
	app.GET("/generateCode", login.GenerateCode)
	//菜单相关
	menurGroup := app.Group("/menu")
	menu := new(controller.MenuController)
	menurGroup.GET("/findMenuList", menu.FindMenuList)
	//用户相关
	userGroup := app.Group("/user")
	user := new(controller.UserController)
	userGroup.POST("/register", user.Register)
	userGroup.POST("/login", user.Login)

}
