package services

import (
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"github.com/AlchemillaHQ/gotham-stack/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"os"
	"testing"
)

var testDb *gorm.DB
var t *testing.T
var TestUserId int
var taskService *TaskService

func setupTestUser() models.User {
	hashed, err := utils.HashPassword("hashedPassword")
	if err != nil {
		t.Fatalf("Hashing password failed: %v", err)
	}

	user := models.User{
		Username: "testuser",
		Email:    "user@example.com",
		Password: hashed,
	}

	if err := testDb.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	TestUserId = user.ID

	return user
}

func teardownTestUser(user models.User) {
	testDb.Unscoped().Delete(&models.User{}, user.ID)
}

func TestMain(m *testing.M) {
	var err error
	testDb, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		panic("failed to connect to database")
	}

	testDb.AutoMigrate(&models.User{}, &models.Task{})

	testUser := setupTestUser()
	defer teardownTestUser(testUser)

	taskService = NewTaskService(testDb)

	code := m.Run()

	os.Exit(code)
}
