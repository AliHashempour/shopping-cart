package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password"`
	Baskets  []*Basket `json:"baskets"  gorm:"foreignKey:UserId"`
}
