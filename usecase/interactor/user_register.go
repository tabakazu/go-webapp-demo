package interactor

import (
	"errors"

	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
)

type userRegister struct {
	Repository domain.UserRepository
}

func NewUserRegister(r domain.UserRepository) userRegister {
	return userRegister{r}
}

func (s userRegister) Execute(user entity.User) (entity.User, error) {
	if user.Password != user.PasswordConfirmation {
		return user, errors.New("password doesn't match password_confirmation")
	}

	passwordDigest, err := user.Password.Digest()
	if err != nil {
		return user, err
	}
	user.PasswordDigest = passwordDigest

	if err := s.Repository.Create(&user); err != nil {
		return user, err
	}

	return user, nil
}
