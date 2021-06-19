package domain

import "github.com/tabakazu/go-webapp/domain/entity"

type UserTokenGenerator interface {
	Issue(*entity.UserAccount) (string, error)
}
