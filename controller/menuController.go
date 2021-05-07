package controller

import (
	"ginWeb/service"
	"ginWeb/util"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
}

var menuService *service.MenuService

func init() {
	menuService = new(service.MenuService)
}
func (this *MenuController) FindMenuList(c *gin.Context) {
	menuList := menuService.FindMenuList()
	util.Success(c, menuList)
}
