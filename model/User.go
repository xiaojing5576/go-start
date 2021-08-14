package model

import (
	"errors"
	"myweb/initsql"
	"time"
)

type User struct {
	Username   string    `gorm:"primary_key;type:varchar(20);"`
	Email      string    `gorm:"type:varchar(20);not null";`
	Sex        string    `gorm:"type:varchar(256);not null;"`
	CreateTime time.Time `gorm:"type:timestamp;not null"`
}

func (user *User) QueryUserWithParam(username string, password string) (User, error) {
	var u User
	initsql.GetDB().Where("username = ? and password = ?", username, password).First(&u)
	return u, errors.New("ok")
}
