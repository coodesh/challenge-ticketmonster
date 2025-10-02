package config

import (
	"fmt"

	"ticketmonster/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Enable UUID extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return nil, fmt.Errorf("failed to create UUID extension: %w", err)
	}

	// Auto-migrate all models
	if err := db.AutoMigrate(
		&models.Movie{},
		&models.User{},
		&models.Showtime{},
		&models.Reservation{},
	); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return db, nil
}
