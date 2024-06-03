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
	Data   string      `json:"data"    gorm:"type:varchar(2048)"`
	State  BasketState `json:"state"   gorm:"type:varchar(255);index;default:PENDING"`
	UserId uint        `json:"user_id"`
}
