package dao

import (
	"ginWeb/models"
)

type MenuDao struct {
}

var table_menu = "t_menu"

func (this *MenuDao) FindMenuList() *[]models.Menu {
	var menu []models.Menu
	mEngine.Table(table_menu).Find(&menu)
	return &menu
}
