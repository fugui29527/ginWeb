package service

import (
	"ginWeb/models"
	"go.uber.org/zap"
)

type MenuService struct {
}

var menu *models.Menu
var logger *zap.Logger

func init() {
	menu = new(models.Menu)
}

func (this *MenuService) FindMenuList() *[]models.Menu {
	return menu.FindMenuList()
}
