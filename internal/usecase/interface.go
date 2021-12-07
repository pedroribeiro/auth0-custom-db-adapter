package usecase

import "github.com/elvenworks/users/internal/domain/entity"

type IUserUseCase interface {
	Create(user *entity.User) error
	Login(email string, password string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}
