package models

import "gorm.io/gorm"

type Count struct {
	gorm.Model
	Count  int
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
