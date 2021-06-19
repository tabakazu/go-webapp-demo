package data

import (
	"github.com/tabakazu/go-webapp/domain/entity"
)

type UserAccountResult struct {
	Username   string `json:"username"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

func NewUserAccountResult(e *entity.UserAccount) *UserAccountResult {
	return &UserAccountResult{
		Username:   e.Username,
		FamilyName: e.FamilyName,
		GivenName:  e.GivenName,
		Email:      e.Email,
	}
}

type LoginResult struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewLoginResult(e *entity.UserAccount, token string) *LoginResult {
	return &LoginResult{
		Username: e.Username,
		Token:    token,
	}
}
