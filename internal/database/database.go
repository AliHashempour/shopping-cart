package database

import (
	"basket/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitializeDB initializes and returns a GORM database connection
func InitializeDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("basket.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate all models
	err = db.AutoMigrate(&model.Basket{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
