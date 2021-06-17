package application

import (
	"context"

	"github.com/tabakazu/go-webapp/domain"
)

type userAccountShowService struct {
	repo domain.UserAccountRepository
}

func NewUserAccountShowService(repo domain.UserAccountRepository) ShowUserAccount {
	return &userAccountShowService{
		repo: repo,
	}
}

func (s *userAccountShowService) Execute(ctx context.Context, userID int) (*ShowUserAccountResult, error) {
	userAccount, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := NewShowUserAccountResult(userAccount)
	return result, nil
}
