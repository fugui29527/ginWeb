package dto

type LoginDto struct {
	Username string `gin:"username" binding:"required" example:"用户名"`
	PassWord string `gin:"password" binding:"required" example:"密码"`
	VailCode string `gin:"vailCode" binding:"required" example:"图形验证码"`
	IdKey    string `gin:"idKey" binding:"required" example:"IdKey"`
}
