package services

import "gorm.io/gorm"

type ServiceRegistry struct {
	AuthService    *AuthService
	CounterService *CounterService
}

func NewServiceRegistry(db *gorm.DB) *ServiceRegistry {
	return &ServiceRegistry{
		AuthService:    NewAuthService(db),
		CounterService: NewCounterService(db),
	}
}
