package application

import (
	"context"

	"github.com/tabakazu/go-webapp/application/data"
)

type RegisterUserAccount interface {
	Execute(context.Context, *data.RegisterUserAccountParam) (*data.RegisterUserAccountResult, error)
}

type LoginUserAccount interface {
	Execute(context.Context, *data.LoginUserAccountParam) (*data.LoginUserAccountResult, error)
}

type ShowUserAccount interface {
	Execute(context.Context, int) (*data.ShowUserAccountResult, error)
}
