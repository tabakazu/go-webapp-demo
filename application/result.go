package application

import (
	"github.com/tabakazu/go-webapp/domain"
)

type RegisterUserAccountResult struct {
	Username   string `json:"username"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

func NewRegisterUserAccountResult(e *domain.UserAccount) *RegisterUserAccountResult {
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

func NewLoginUserAccountResult(e *domain.UserAccount, token string) *LoginUserAccountResult {
	return &LoginUserAccountResult{
		Username: e.Username,
		Token:    token,
	}
}
