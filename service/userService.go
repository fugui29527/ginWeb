package service

import (
	"ginWeb/dao"
	"ginWeb/models"
)

type UserService struct {
}

func (this *UserService) Insert(u *models.UserInfo) (int64, error) {
	userDao := new(dao.UserInfoDao)
	return userDao.Insert(u)
}
