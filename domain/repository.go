package domain

import "context"

type UserAccountRepository interface {
	Create(context.Context, *UserAccount) error
}
