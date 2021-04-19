package usecase

import (
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type UserRegister interface {
	Execute(entity.User) (entity.User, error)
}

type UserAuthenticate interface {
	Execute(value.Email, value.Password) (entity.User, error)
}
