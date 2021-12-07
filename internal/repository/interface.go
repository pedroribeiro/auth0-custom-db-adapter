package repository

import "github.com/elvenworks/users/internal/domain/entity"

type IUserRepo interface {
	Create(user *entity.User) error
	Get(email string) (*entity.User, error)
}
