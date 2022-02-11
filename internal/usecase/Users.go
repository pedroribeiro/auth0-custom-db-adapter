package usecase

import (
	"errors"
	"fmt"

	"github.com/pedroribeiro/users/internal/domain/entity"
	"github.com/pedroribeiro/users/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	Repo repository.IUserRepo
}

func (u *UserUseCase) Create(user *entity.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hash)

	if err != nil {
		return err
	}

	if err = u.Repo.Create(user); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) Login(email string, password string) (*entity.User, error) {

	user, err := u.Repo.Get(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println("eeeerrrjl", err)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func (u *UserUseCase) GetByEmail(email string) (*entity.User, error) {

	user, err := u.Repo.Get(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
