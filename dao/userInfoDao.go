package dao

import (
	"ginWeb/models"
	"ginWeb/util"
)

type UserInfoDao struct {
}

func (this *UserInfoDao) Insert(u *models.UserInfo) (int64, error) {
	result, err := util.DbEngin.Insert(u)
	return result, err
}
