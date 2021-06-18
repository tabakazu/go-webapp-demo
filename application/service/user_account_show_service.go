package service

import (
	"context"

	"github.com/tabakazu/go-webapp/application"
	"github.com/tabakazu/go-webapp/application/data"
	"github.com/tabakazu/go-webapp/domain"
)

type userAccountShowService struct {
	repo domain.UserAccountRepository
}

func NewUserAccountShowService(repo domain.UserAccountRepository) application.ShowUserAccount {
	return &userAccountShowService{
		repo: repo,
	}
}

func (s *userAccountShowService) Execute(ctx context.Context, userID int) (*data.ShowUserAccountResult, error) {
	userAccount, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := data.NewShowUserAccountResult(userAccount)
	return result, nil
}
