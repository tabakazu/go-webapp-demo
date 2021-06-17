package domain

import "context"

type UserAccountRepository interface {
	FindByID(context.Context, int) (*UserAccount, error)
	FindByUsername(context.Context, string) (*UserAccount, error)
	FindByEmail(context.Context, string) (*UserAccount, error)
	Create(context.Context, *UserAccount) error
}
