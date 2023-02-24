package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int64  `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	FullName string `json:"fullname" gorm:"not null"`
	Avatar   string `json:"avatar" gorm:"not null;unique index"`
}
