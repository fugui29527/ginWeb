package models

type AdminUser struct {
	Id        int64  `xorm:"pk autoincr" json:"id"`
	Name      string `xorm:"name" json:"name" `
	Phone     string `xorm:"phone" json:"phone" `
	Telephone string `xorm:"telephone" json:"telephone" `
	Address   string `xorm:"address" json:"address" `
	Enabled   int8   `xorm:"enabled" json:"enabled" `
	Username  string `xorm:"username" json:"enabled" `
	PassWord  string `xorm:"passWord" json:"passWord" `
	UserFace  string `xorm:"userFace" json:"userFace" `
	Remark    string `xorm:"remark" json:"remark" `
}

var table_admin = "t_admin"

func CreateUserInfo(name string,
	phone string,
	telephone string,
	address string,
	enabled int8,
	username string,
	passWord string,
	userFace string,
	remark string) *AdminUser {

	return &AdminUser{
		Name:      name,
		Phone:     phone,
		Telephone: telephone,
		Address:   address,
		Enabled:   enabled,
		Username:  username,
		PassWord:  passWord,
		UserFace:  userFace,
		Remark:    remark,
	}
}

func (this *AdminUser) Insert(u *AdminUser) (int64, error) {
	result, err := mEngine.Insert(u)
	return result, err
}

//查找用户
func (this *AdminUser) FindUser() (bool, error) {
	has, err := mEngine.Table(table_admin).Get(this)
	if err == nil && has {
		return true, nil
	}

	return false, err
}
