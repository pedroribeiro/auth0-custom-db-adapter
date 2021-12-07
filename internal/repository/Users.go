package repository

import (
	"github.com/elvenworks/users/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) Create(user *entity.User) error {
	tx := r.DB.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *UserRepo) Get(email string) (*entity.User, error) {
	var user entity.User
	tx := r.DB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
