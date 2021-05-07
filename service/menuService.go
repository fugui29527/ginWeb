package service

import (
	"ginWeb/dao"
	"ginWeb/models"
	"go.uber.org/zap"
)

type MenuService struct {
}

var menudao *dao.MenuDao
var logger *zap.Logger

func init() {
	menudao = new(dao.MenuDao)
}

func (this *MenuService) FindMenuList() *[]models.Menu {
	return menudao.FindMenuList()
}
