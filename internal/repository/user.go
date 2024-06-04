package repository

import (
	"basket/internal/model"
	"gorm.io/gorm"
)

type User interface {
	GetBy(query interface{}, args ...interface{}) (*model.User, error)
	Create(user *model.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) GetBy(query interface{}, args ...interface{}) (*model.User, error) {
	var user model.User
	err := repo.db.Where(query, args...).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) Create(user *model.User) error {
	return repo.db.Create(user).Error
}
