package dao

import (
	"ginWeb/models"
)

type UserInfoDao struct {
}

func (this *UserInfoDao) Insert(u *models.UserInfo) (int64, error) {
	result, err := mEngine.Insert(u)
	return result, err
}
