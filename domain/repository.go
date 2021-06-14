package domain

import "context"

type UserAccountRepository interface {
	FindByUsername(context.Context, string) (*UserAccount, error)
	FindByEmail(context.Context, string) (*UserAccount, error)
	Create(context.Context, *UserAccount) error
}
