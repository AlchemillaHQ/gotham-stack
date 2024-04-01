package services

import (
	"errors"
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"github.com/AlchemillaHQ/gotham-stack/internal/utils"
	"gorm.io/gorm"
	"testing"
)

func setupTestUser(t *testing.T) models.User {
	hashed, err := utils.HashPassword("hashedPassword")

	if err != nil {
		t.Fatalf("Hashing password failed: %v", err)
	}

	user := models.User{
		Username: "testuser",
		Email:    "user@example.com",
		Password: hashed,
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	return user
}

func teardownTestUser(t *testing.T, user models.User) {
	db.Unscoped().Delete(&models.User{}, user.ID)
}

func TestIncrementCount(t *testing.T) {
	user := setupTestUser(t)
	defer teardownTestUser(t, user)

	counterService := NewCounterService(db)

	newCount, err := counterService.IncrementCount(int(user.ID))
	if err != nil {
		t.Errorf("IncrementCount() error = %v, wantErr %v", err, false)
	}
	if newCount != 1 {
		t.Errorf("Expected new count to be 1, got %d", newCount)
	}

	newCount, err = counterService.IncrementCount(int(user.ID))
	if err != nil {
		t.Errorf("IncrementCount() error = %v, wantErr %v", err, false)
	}
	if newCount != 2 {
		t.Errorf("Expected new count to be 2, got %d", newCount)
	}
}

func TestDecrementCount(t *testing.T) {
	user := setupTestUser(t)
	defer teardownTestUser(t, user)

	counterService := NewCounterService(db)

	newCount, err := counterService.DecrementCount(int(user.ID))
	if err != nil {
		t.Errorf("DecrementCount() error = %v, wantErr %v", err, false)
	}
	if newCount != -1 {
		t.Errorf("Expected new count to be -1, got %d", newCount)
	}

	newCount, err = counterService.DecrementCount(int(user.ID))
	if err != nil {
		t.Errorf("DecrementCount() error = %v, wantErr %v", err, false)
	}
	if newCount != -2 {
		t.Errorf("Expected new count to be -2, got %d", newCount)
	}
}

func TestGetCount(t *testing.T) {
	user := setupTestUser(t)
	defer teardownTestUser(t, user)

	counterService := NewCounterService(db)

	initialCount, err := counterService.GetCount(int(user.ID))

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("GetCount() error = %v, wantErr %v", err, false)
	}

	if initialCount != 0 {
		t.Errorf("Expected initial count to be 0, got %d", initialCount)
	}

	_, err = counterService.IncrementCount(int(user.ID))
	if err != nil {
		t.Fatalf("IncrementCount() error = %v, wantErr %v", err, false)
	}

	countAfterIncrement, err := counterService.GetCount(int(user.ID))
	if err != nil {
		t.Fatalf("GetCount() after increment error = %v, wantErr %v", err, false)
	}
	if countAfterIncrement != 1 {
		t.Errorf("Expected count after increment to be 1, got %d", countAfterIncrement)
	}

	_, err = counterService.DecrementCount(int(user.ID))
	if err != nil {
		t.Fatalf("DecrementCount() error = %v, wantErr %v", err, false)
	}

	finalCount, err := counterService.GetCount(int(user.ID))
	if err != nil {
		t.Fatalf("GetCount() after decrement error = %v, wantErr %v", err, false)
	}
	if finalCount != 0 {
		t.Errorf("Expected final count to be 0, got %d", finalCount)
	}
}
