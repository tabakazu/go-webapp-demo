package application

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type RegisterUserAccountParam struct {
	Username             string `json:"username"`
	FamilyName           string `json:"family_name"`
	GivenName            string `json:"given_name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func ValidateRegisterUserAccountParam(p *RegisterUserAccountParam) error {
	validate = validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}

type LoginUserAccountParam struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

func ValidateLoginUserAccountParam(p *LoginUserAccountParam) error {
	validate = validator.New()
	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}
