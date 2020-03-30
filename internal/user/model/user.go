package model

import (
	"time"
)

type UserLoginParam struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type UserRegisterParam struct {
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
	NickName string `json:"nickname"`
}

func (ur *UserRegisterParam) ToUser() *User {
	return &User{
		Email:    ur.Email,
		Pwd:      ur.Pwd,
		NickName: ur.NickName,
	}
}

type User struct {
	Uid       uint64    `gorm:"primary_key;auto_increment" json:"uid"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Pwd       string    `gorm:"size:100;not null;" json:"pwd"`
	NickName  string    `gorm:"size:255;not null" json:"nickname"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
