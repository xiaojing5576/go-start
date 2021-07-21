package model

import "time"

type User struct {
	Username   string    `gorm:"primary_key;type:varchar(20);"`
	Email      string    `gorm:"type:varchar(20);not null";`
	Sex        string    `gorm:"type:varchar(256);not null;"`
	CreateTime time.Time `gorm:"type:timestamp;not null"`
}
