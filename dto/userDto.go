package dto

type UserRegister struct {
	Username string `gin:"username"`
	PassWord string `gin:"password"`
	Sex      int    `gin:"sex"`
}
