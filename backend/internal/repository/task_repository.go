// Package repository provides data access functions for the domain models.
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"backend/internal/models"
)

// ErrTaskNotFound is returned when a requested task does not exist.
var ErrTaskNotFound = errors.New("task not found")

// TaskRepository defines the persistence operations available for tasks.
type TaskRepository interface {
	Create(task *models.Task) error
	FindAll() ([]models.Task, error)
	FindByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a TaskRepository backed by the given database connection.
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (r *taskRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Order("id asc").Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}
	return tasks, nil
}

func (r *taskRepository) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, fmt.Errorf("failed to fetch task: %w", err)
	}
	return &task, nil
}

func (r *taskRepository) Update(task *models.Task) error {
	result := r.db.Model(task).Clauses().Select(
		"Title", "Description", "Status", "Priority", "AssignedTo", "DueDate",
	).Updates(task)
	if result.Error != nil {
		return fmt.Errorf("failed to update task: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}
	return nil
}

func (r *taskRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Task{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete task: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}
	return nil
}
