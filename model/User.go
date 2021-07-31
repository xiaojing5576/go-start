package model

import (
	"myweb/initsql"
	"time"
)

type User struct {
	Username   string    `gorm:"primary_key;type:varchar(20);"`
	Email      string    `gorm:"type:varchar(20);not null";`
	Sex        string    `gorm:"type:varchar(256);not null;"`
	CreateTime time.Time `gorm:"type:timestamp;not null"`
}

func QueryUserWithParam(username string, password string) User {
	var user User
	initsql.GetDB().Where("username = ? and password = ?", username, password).First(&user)
	return user
}
