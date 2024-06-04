package repository

import (
	"basket/internal/model"
	"gorm.io/gorm"
)

// Basket interface
type Basket interface {
	Get(basketId *uint, userId *uint) ([]*model.Basket, error)
	Create(*model.Basket) error
	Update(*model.Basket) error
	Delete(*model.Basket) error
}

// BasketRepo struct
type BasketRepo struct {
	db *gorm.DB
}

// NewBasketRepo creates a new BasketRepo
func NewBasketRepo(db *gorm.DB) *BasketRepo {
	return &BasketRepo{db: db}
}

// Get retrieves one or more baskets based on the provided ID
func (repo *BasketRepo) Get(basketId *uint, userId *uint) ([]*model.Basket, error) {
	var baskets []*model.Basket
	query := repo.db

	if basketId != nil {
		query = query.Where("id = ?", basketId)
	}
	if userId != nil {
		query = query.Where("user_id = ?", userId)
	}

	if err := query.Find(&baskets).Error; err != nil {
		return nil, err
	}
	return baskets, nil
}

// Create creates a new basket
func (repo *BasketRepo) Create(basket *model.Basket) error {
	return repo.db.Create(basket).Error
}

// Update updates an existing basket
func (repo *BasketRepo) Update(basket *model.Basket) error {
	return repo.db.Save(basket).Error
}

// Delete deletes a basket
func (repo *BasketRepo) Delete(basket *model.Basket) error {
	return repo.db.Delete(basket).Error
}
