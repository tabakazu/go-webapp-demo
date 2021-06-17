package application

import (
	"context"
)

type RegisterUserAccount interface {
	Execute(context.Context, *RegisterUserAccountParam) (*RegisterUserAccountResult, error)
}

type LoginUserAccount interface {
	Execute(context.Context, *LoginUserAccountParam) (*LoginUserAccountResult, error)
}

type ShowUserAccount interface {
	Execute(context.Context, int) (*ShowUserAccountResult, error)
}
