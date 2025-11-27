package db

import (
	"fmt"
	"gin-quickstart/internal/albums"
	"gin-quickstart/internal/auth"
	"gin-quickstart/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the GORM database connection and runs migrations.
func InitDB(cfg config.Config) (*gorm.DB, error) {
	// Construct the PostgreSQL DSN (Data Source Name) string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.SSLMode,
	)
	// Open and connect to the PostgreSQL database using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//  Check for connection errors.
	if err != nil {
		return nil, err
	}

	// Run AutoMigrate for the models.Album struct.
	if err := db.AutoMigrate(
		&albums.Album{},
		&auth.User{},
	); err != nil {
		return nil, err
	}

	// Return the *gorm.DB instance and nil error, or nil and the error.
	return db, nil
}
