package application

import (
	"github.com/tabakazu/go-webapp/domain"
)

type RegisterUserAccountParam struct {
	Username             string `json:"username"`
	FamilyName           string `json:"family_name"`
	GivenName            string `json:"given_name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

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

type LoginUserAccountParam struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

type LoginUserAccountResult struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewLoginUserAccountResult(e *domain.UserAccount) *LoginUserAccountResult {
	return &LoginUserAccountResult{
		ID:       e.ID,
		Username: e.Username,
	}
}
