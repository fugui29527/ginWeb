package main

import (
	"ginWeb/controller"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func initRouter(app *gin.Engine) {
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//登录相关
	login := new(controller.LoginController)
	app.GET("/generateCode", login.GenerateCode)
	app.POST("/login", login.Login)
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
