package service

import (
	"ginWeb/models"
)

type UserService struct {
}

func (this *UserService) Insert(u *models.AdminUser) (int64, error) {
	adminUser := new(models.AdminUser)
	return adminUser.Insert(u)
}

func (this *UserService) FindUser(userName string) (*models.AdminUser, error) {
	u := &models.AdminUser{
		Username: userName,
	}
	has, err := u.FindUser()
	if err == nil && has {
		return u, nil
	}
	return nil, err
}
