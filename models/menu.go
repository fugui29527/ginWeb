package models

type Menu struct {
	Id          int64  `json:"id" xorm:"not null pk autoincr"`
	Url         string `json:"url" xorm:"url"`
	Path        string `json:"path" xorm:"path"`
	Component   string `json:"component" xorm:"component"`
	Name        string `json:"name" xorm:"name"`
	IconCls     string `json:"iconCls" xorm:"iconCls"`
	KeepAlive   int8   `json:"keepAlive" xorm:"keepAlive"`
	RequireAuth int8   `json:"requireAuth" xorm:"requireAuth"`
	ParentId    int64  `json:"parentId" xorm:"parentId"`
	Enabled     int8   `json:"enabled" xorm:"enabled"`
}

func CreateMenu(url string,
	path string,
	component string,
	name string,
	iconCls string,
	keepAlive int8,
	requireAuth int8,
	parentId int64,
	enabled int8) *Menu {
	return &Menu{
		Url:         url,
		Path:        path,
		Component:   component,
		Name:        name,
		IconCls:     iconCls,
		KeepAlive:   keepAlive,
		RequireAuth: requireAuth,
		ParentId:    parentId,
		Enabled:     enabled,
	}
}

var table_menu = "t_menu"

func (this *Menu) FindMenuList() *[]Menu {
	var menu []Menu
	mEngine.Table(table_menu).Find(&menu)
	return &menu
}
