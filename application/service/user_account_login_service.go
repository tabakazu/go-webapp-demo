package service

import (
	"context"

	"github.com/tabakazu/go-webapp/application"
	"github.com/tabakazu/go-webapp/application/data"
	"github.com/tabakazu/go-webapp/domain"
	"github.com/tabakazu/go-webapp/domain/entity"
)

type userAccountLoginService struct {
	repo domain.UserAccountRepository
	gen  domain.UserTokenGenerator
}

func NewUserAccountLoginService(
	repo domain.UserAccountRepository,
	gen domain.UserTokenGenerator,
) application.LoginUserAccount {
	return &userAccountLoginService{
		repo: repo,
		gen:  gen,
	}
}

func (s *userAccountLoginService) Execute(ctx context.Context, param *data.LoginUserAccountParam) (*data.LoginResult, error) {
	userAccount, err := func() (*entity.UserAccount, error) {
		if param.UsernameOrEmail.IsEmail() {
			return s.repo.FindByEmail(ctx, string(param.UsernameOrEmail))
		} else {
			return s.repo.FindByUsername(ctx, string(param.UsernameOrEmail))
		}
	}()
	if err != nil {
		return nil, err
	}

	if err := userAccount.ValidPassword(param.Password); err != nil {
		return nil, err
	}

	token, err := s.gen.Issue(userAccount)
	if err != nil {
		return nil, err
	}

	result := data.NewLoginResult(userAccount, token)
	return result, nil
}
