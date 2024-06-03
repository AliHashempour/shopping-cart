package repository

import (
	"basket/internal/model"
	"gorm.io/gorm"
)

type User interface {
	Get(*uint) error
	Create(*model.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Get(id *uint) error {
	return repo.db.First(&model.User{}, id).Error
}

func (repo *UserRepo) Create(user *model.User) error {
	return repo.db.Create(user).Error
}
