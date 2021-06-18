package data

import (
	"github.com/tabakazu/go-webapp/domain/entity"
)

type RegisterUserAccountResult struct {
	Username   string `json:"username"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

func NewRegisterUserAccountResult(e *entity.UserAccount) *RegisterUserAccountResult {
	return &RegisterUserAccountResult{
		Username:   e.Username,
		FamilyName: e.FamilyName,
		GivenName:  e.GivenName,
		Email:      e.Email,
	}
}

type LoginUserAccountResult struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewLoginUserAccountResult(e *entity.UserAccount, token string) *LoginUserAccountResult {
	return &LoginUserAccountResult{
		Username: e.Username,
		Token:    token,
	}
}

type ShowUserAccountResult struct {
	Username   string `json:"username"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

func NewShowUserAccountResult(e *entity.UserAccount) *ShowUserAccountResult {
	return &ShowUserAccountResult{
		Username:   e.Username,
		FamilyName: e.FamilyName,
		GivenName:  e.GivenName,
		Email:      e.Email,
	}
}
