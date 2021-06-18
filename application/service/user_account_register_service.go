package service

import (
	"context"
	"errors"

	"github.com/tabakazu/go-webapp/application"
	"github.com/tabakazu/go-webapp/application/data"
	"github.com/tabakazu/go-webapp/domain"
	"github.com/tabakazu/go-webapp/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type userAccountRegisterService struct {
	repo domain.UserAccountRepository
}

func NewUserAccountRegisterService(repo domain.UserAccountRepository) application.RegisterUserAccount {
	return &userAccountRegisterService{
		repo: repo,
	}
}

func (s *userAccountRegisterService) Execute(ctx context.Context, param *data.RegisterUserAccountParam) (*data.RegisterUserAccountResult, error) {
	if param.Password != param.PasswordConfirmation {
		return nil, errors.New("password doesn't match password_confirmation")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userAccount := entity.UserAccount{
		Username:       param.Username,
		FamilyName:     param.FamilyName,
		GivenName:      param.GivenName,
		Email:          param.Email,
		PasswordDigest: string(passHash),
	}
	if err := s.repo.Create(ctx, &userAccount); err != nil {
		return nil, err
	}

	result := data.NewRegisterUserAccountResult(&userAccount)
	return result, nil
}
