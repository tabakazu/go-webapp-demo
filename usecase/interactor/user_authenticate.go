package interactor

import (
	"errors"

	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type userAuthenticate struct {
	Repository domain.UserRepository
}

func NewUserAuthenticate(r domain.UserRepository) userAuthenticate {
	return userAuthenticate{r}
}

func (s userAuthenticate) Execute(email value.Email, pass value.Password) (entity.User, error) {
	user, err := s.Repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	passwordDigest, err := pass.Digest()
	if err != nil {
		return user, err
	}

	if user.PasswordDigest == passwordDigest {
		return user, errors.New("Invalid password")
	}

	return user, nil
}
