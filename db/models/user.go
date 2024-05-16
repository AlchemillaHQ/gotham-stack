package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement;"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}
