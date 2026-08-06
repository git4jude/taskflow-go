// Package models defines the domain entities persisted by the application.
package models

import "time"

// Task represents a unit of work tracked by the system.
type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"not null" binding:"required"`
	Description string     `json:"description"`
	Status      string     `json:"status" gorm:"not null;default:todo" binding:"omitempty,oneof=todo in_progress done"`
	Priority    string     `json:"priority" gorm:"not null;default:medium" binding:"omitempty,oneof=low medium high"`
	AssignedTo  string     `json:"assigned_to"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
