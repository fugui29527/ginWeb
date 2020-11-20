package models

import "time"

type UserInfo struct {
	Id         int64     `xorm:"pk autoincr" json:"id"`
	UserName   string    `xorm:"userName" json:"userName" `
	PassWord   string    `xorm:"passWord" json:"passWord" `
	Sex        int       `xorm:"sex" json:"sex" `
	Phone      string    `xorm:"phone" json:"phone"`
	CreateTime time.Time `xorm:"create_time" json:"createTime"`
}

func CreateUserInfo(userName string, password string, sex int) *UserInfo {

	return &UserInfo{
		UserName:   userName,
		PassWord:   password,
		Sex:        sex,
		CreateTime: time.Now(),
	}
}
