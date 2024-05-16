package services

import (
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"github.com/AlchemillaHQ/gotham-stack/internal/utils"
	"testing"
)

func TestVerifyCredentials(t *testing.T) {
	password := "securepassword"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Fatalf("Hashing password failed: %v", err)
	}
	user := models.User{
		Email:    "test@example.com",
		Password: hashedPassword,
	}
	if err := testDb.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	authService := NewAuthService(testDb)

	testCases := []struct {
		name          string
		email         string
		password      string
		expectedValid bool
	}{
		{"Correct credentials", "test@example.com", "securepassword", true},
		{"Incorrect password", "test@example.com", "wrongpassword", false},
		{"Non-existent user", "nouser@example.com", "securepassword", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, valid := authService.VerifyCredentials(tc.email, tc.password)
			if valid != tc.expectedValid {
				t.Errorf("Expected validity: %v, got: %v for scenario: %s", tc.expectedValid, valid, tc.name)
			}
		})
	}

	testDb.Where("email = ?", user.Email).Delete(&models.User{})
}
