package application

import (
	"context"
)

type RegisterUserAccount interface {
	Execute(context.Context, *RegisterUserAccountParam) (*RegisterUserAccountResult, error)
}
