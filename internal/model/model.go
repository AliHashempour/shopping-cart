package model

import (
	"gorm.io/gorm"
)

// BasketState defines the type for the state of Basket
type BasketState string

const (
	StateCompleted BasketState = "COMPLETED"
	StatePending   BasketState = "PENDING"
)

// Basket represents a basket in the database.
type Basket struct {
	gorm.Model
	Data  string      `gorm:"type:varchar(2048)"`                      // Ensures the data does not exceed 2048 bytes
	State BasketState `gorm:"type:varchar(255);index;default:PENDING"` // Indexed for faster queries on state, default to PENDING
}
