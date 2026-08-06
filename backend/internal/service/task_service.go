// Package service implements the business logic for managing tasks.
package service

import (
	"fmt"

	"backend/internal/models"
	"backend/internal/repository"
)

// TaskService defines the business operations available for tasks.
type TaskService interface {
	CreateTask(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(id uint, input *models.Task) (*models.Task, error)
	DeleteTask(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a TaskService backed by the given repository.
func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	if task.Status == "" {
		task.Status = "todo"
	}
	if task.Priority == "" {
		task.Priority = "medium"
	}

	if err := s.repo.Create(task); err != nil {
		return fmt.Errorf("service: create task: %w", err)
	}
	return nil
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	tasks, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("service: get all tasks: %w", err)
	}
	return tasks, nil
}

func (s *taskService) GetTaskByID(id uint) (*models.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) UpdateTask(id uint, input *models.Task) (*models.Task, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existing.Title = input.Title
	existing.Description = input.Description
	existing.AssignedTo = input.AssignedTo
	existing.DueDate = input.DueDate

	if input.Status != "" {
		existing.Status = input.Status
	}
	if input.Priority != "" {
		existing.Priority = input.Priority
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, fmt.Errorf("service: update task: %w", err)
	}
	return existing, nil
}

func (s *taskService) DeleteTask(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
