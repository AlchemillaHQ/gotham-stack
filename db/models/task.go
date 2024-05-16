package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID        int `gorm:"primaryKey;autoIncrement;"`
	UserID    int
	Task      string
	CreatedAt time.Time
}
