package domain

import (
	"context"

	"github.com/tabakazu/go-webapp/domain/entity"
)

type UserAccountRepository interface {
	FindByID(context.Context, int) (*entity.UserAccount, error)
	FindByUsername(context.Context, string) (*entity.UserAccount, error)
	FindByEmail(context.Context, string) (*entity.UserAccount, error)
	Create(context.Context, *entity.UserAccount) error
}
