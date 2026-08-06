// Package handler implements the HTTP layer, translating requests to service calls.
package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/service"
)

// TaskHandler exposes HTTP handlers for task resources.
type TaskHandler struct {
	service service.TaskService
}

// NewTaskHandler creates a TaskHandler backed by the given service.
func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// taskInput is the request body accepted for creating and updating tasks.
type taskInput struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Status      string  `json:"status" binding:"omitempty,oneof=todo in_progress done"`
	Priority    string  `json:"priority" binding:"omitempty,oneof=low medium high"`
	AssignedTo  string  `json:"assigned_to"`
	DueDate     *string `json:"due_date"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// Create handles POST /api/tasks.
func (h *TaskHandler) Create(c *gin.Context) {
	var input taskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	task, err := toModel(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	if err := h.service.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// List handles GET /api/tasks.
func (h *TaskHandler) List(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: "failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Get handles GET /api/tasks/:id.
func (h *TaskHandler) Get(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "invalid task id"})
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		respondTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

// Update handles PUT /api/tasks/:id.
func (h *TaskHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "invalid task id"})
		return
	}

	var input taskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	task, err := toModel(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	updated, err := h.service.UpdateTask(id, task)
	if err != nil {
		respondTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, updated)
}

// Delete handles DELETE /api/tasks/:id.
func (h *TaskHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "invalid task id"})
		return
	}

	if err := h.service.DeleteTask(id); err != nil {
		respondTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

// Health handles GET /api/health.
func (h *TaskHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func respondTaskError(c *gin.Context, err error) {
	if errors.Is(err, repository.ErrTaskNotFound) {
		c.JSON(http.StatusNotFound, errorResponse{Error: "task not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, errorResponse{Error: "internal server error"})
}

func toModel(input taskInput) (*models.Task, error) {
	task := &models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Priority:    input.Priority,
		AssignedTo:  input.AssignedTo,
	}

	if input.DueDate != nil && *input.DueDate != "" {
		dueDate, err := parseDueDate(*input.DueDate)
		if err != nil {
			return nil, err
		}
		task.DueDate = dueDate
	}

	return task, nil
}

// parseDueDate accepts either RFC3339 timestamps or plain "YYYY-MM-DD" dates.
func parseDueDate(value string) (*time.Time, error) {
	if t, err := time.Parse(time.RFC3339, value); err == nil {
		return &t, nil
	}
	if t, err := time.Parse("2006-01-02", value); err == nil {
		return &t, nil
	}
	return nil, fmt.Errorf("due_date must be RFC3339 or YYYY-MM-DD format")
}
