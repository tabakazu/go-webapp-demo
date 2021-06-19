package application

import (
	"context"

	"github.com/tabakazu/go-webapp/application/data"
)

type RegisterUserAccount interface {
	Execute(context.Context, *data.RegisterUserAccountParam) (*data.UserAccountResult, error)
}

type LoginUserAccount interface {
	Execute(context.Context, *data.LoginUserAccountParam) (*data.LoginResult, error)
}

type ShowUserAccount interface {
	Execute(context.Context, int) (*data.UserAccountResult, error)
}
