package data

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type RegisterUserAccountParam struct {
	Username             string `json:"username" validate:"required"`
	FamilyName           string `json:"family_name" validate:"required"`
	GivenName            string `json:"given_name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

func ValidateRegisterUserAccountParam(p *RegisterUserAccountParam) error {
	validate = validator.New()
	validate.RegisterStructValidation(registerUserAccountParamStructLevelValidation, RegisterUserAccountParam{})

	if err := validate.Struct(p); err != nil {
		return err
	}
	return nil
}

func registerUserAccountParamStructLevelValidation(sl validator.StructLevel) {
	param := sl.Current().Interface().(RegisterUserAccountParam)

	if param.Password != param.PasswordConfirmation {
		sl.ReportError(param.PasswordConfirmation, "PasswordConfirmation", "PasswordConfirmation", "password_confirmation", "")
	}
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
