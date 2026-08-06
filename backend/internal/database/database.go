// Package database manages the PostgreSQL connection and schema migrations.
package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"backend/internal/models"
)

// Connect opens a GORM connection to PostgreSQL using the given DSN.
func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// AutoMigrate creates or updates database tables to match the current models.
func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return fmt.Errorf("failed to auto-migrate database: %w", err)
	}
	return nil
}
