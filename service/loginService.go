package service

import (
	"fmt"
	"ginWeb/dto"
	"ginWeb/errors"
	"ginWeb/models"
	"ginWeb/util"
)

type LoginService struct {
}

func (this *LoginService) DoLogin(loginDto *dto.LoginDto) (token string, err error) {
	//查找用户是否存在
	u := &models.AdminUser{
		Username: loginDto.Username,
	}
	has, err := u.FindUser()
	if err != nil {
		return "", err
	}
	if !has {
		noErr := errors.New("没有此用户")
		return "", noErr
	}
	//验证登录密码
	fmt.Println(util.GetMd5String(loginDto.PassWord))
	if util.GetMd5String(loginDto.PassWord) != u.PassWord {
		err := errors.New("密码错误")
		return "", err
	}
	//生成token
	token, tokenErr := util.CreateToken(&util.Claims{
		UserId:   u.Id,
		Username: u.Username,
	})
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}
