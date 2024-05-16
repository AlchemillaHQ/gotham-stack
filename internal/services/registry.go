package services

import "gorm.io/gorm"

type ServiceRegistry struct {
	AuthService *AuthService
	TaskService *TaskService
}

func NewServiceRegistry(db *gorm.DB) *ServiceRegistry {
	return &ServiceRegistry{
		AuthService: NewAuthService(db),
		TaskService: NewTaskService(db),
	}
}
