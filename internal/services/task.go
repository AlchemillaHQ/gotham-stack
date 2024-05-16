package services

import (
	"github.com/AlchemillaHQ/gotham-stack/db/models"
	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{DB: db}
}

func (service *TaskService) GetTask(userID int, taskID int) (*models.Task, error) {
	var task models.Task
	result := service.DB.Where("user_id = ? AND id = ?", userID, taskID).First(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (service *TaskService) GetTasks(userID int) ([]models.Task, error) {
	var tasks []models.Task
	result := service.DB.Where("user_id = ?", userID).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (service *TaskService) CreateTask(userID int, task string) (int, error) {
	newTask := models.Task{
		UserID: userID,
		Task:   task,
	}

	result := service.DB.Create(&newTask)

	return newTask.ID, result.Error
}

func (service *TaskService) DeleteTask(userID int, taskID int) error {
	result := service.DB.Where("user_id = ? AND id = ?", userID, taskID).Delete(&models.Task{})

	return result.Error
}

func (service *TaskService) DeleteTasks(userID int) error {
	result := service.DB.Where("user_id = ?", userID).Delete(&models.Task{})

	return result.Error
}

func (service *TaskService) GetTaskCount(userID int) (int64, error) {
	var count int64
	result := service.DB.Model(&models.Task{}).Where("user_id = ?", userID).Count(&count)

	return count, result.Error
}
