package services

import (
	"errors"
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"gorm.io/gorm"
)

type CounterService struct {
	DB *gorm.DB
}

func NewCounterService(db *gorm.DB) *CounterService {
	return &CounterService{DB: db}
}

func (service *CounterService) GetCount(userID int) (int, error) {
	var count models.Count
	result := service.DB.Where("user_id = ?", userID).First(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return count.Count, nil
}

func (service *CounterService) DecrementCount(userID int) (int, error) {
	var count models.Count
	result := service.DB.Where("user_id = ?", userID).First(&count)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		count = models.Count{
			UserID: uint(userID),
			Count:  0,
		}
		result = service.DB.Create(&count)
		if result.Error != nil {
			return 0, result.Error
		}
	} else if result.Error != nil {
		return 0, result.Error
	}

	count.Count--
	result = service.DB.Save(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count.Count, nil
}

func (service *CounterService) IncrementCount(userID int) (int, error) {
	var count models.Count
	result := service.DB.Where("user_id = ?", userID).First(&count)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		count = models.Count{
			UserID: uint(userID),
			Count:  0,
		}
		result = service.DB.Create(&count)
		if result.Error != nil {
			return 0, result.Error
		}
	} else if result.Error != nil {
		return 0, result.Error
	}

	count.Count++
	result = service.DB.Save(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count.Count, nil
}
