package services

import (
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Count{})

	code := m.Run()

	os.Exit(code)
}
