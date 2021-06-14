package application

import (
	"context"

	"github.com/tabakazu/go-webapp/domain"
	"golang.org/x/crypto/bcrypt"
)

type userAccountLoginService struct {
	repo domain.UserAccountRepository
}

func NewUserAccountLoginService(repo domain.UserAccountRepository) LoginUserAccount {
	return &userAccountLoginService{
		repo: repo,
	}
}

func (s *userAccountLoginService) Execute(ctx context.Context, param *LoginUserAccountParam) (*LoginUserAccountResult, error) {
	userAccount, err := s.repo.FindByUsername(ctx, param.UsernameOrEmail)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userAccount.PasswordDigest), []byte(param.Password)); err != nil {
		return nil, err
	}

	result := NewLoginUserAccountResult(userAccount)
	return result, nil
}
