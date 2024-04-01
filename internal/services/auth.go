package services

import (
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"github.com/AlchemillaHQ/gotham-stack/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (service *AuthService) VerifyCredentials(email, password string) (*models.User, bool) {
	var user models.User
	result := service.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, false
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, false
	}

	return &user, true
}
