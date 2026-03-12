package repository

import (
	"task-service/config"
	"task-service/models"

	"github.com/google/uuid"
)

func GetAllTasks(status, priority string) ([]models.Task, error) {
	var tasks []models.Task
	query := config.DB
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}
	result := query.Find(&tasks)
	return tasks, result.Error
}

func GetTaskByID(id uuid.UUID) (models.Task, error) {
	var task models.Task
	result := config.DB.First(&task, "id = ?", id)
	return task, result.Error
}

func CreateTask(task *models.Task) error {
	return config.DB.Create(task).Error
}

func UpdateTask(task *models.Task) error {
	return config.DB.Save(task).Error
}

func DeleteTask(id uuid.UUID) error {
	return config.DB.Delete(&models.Task{}, "id = ?", id).Error
}
