package service

import (
	"context"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tabakazu/go-webapp/application"
	"github.com/tabakazu/go-webapp/application/data"
	"github.com/tabakazu/go-webapp/domain"
	"golang.org/x/crypto/bcrypt"
)

type userAccountLoginService struct {
	repo domain.UserAccountRepository
}

func NewUserAccountLoginService(repo domain.UserAccountRepository) application.LoginUserAccount {
	return &userAccountLoginService{
		repo: repo,
	}
}

func (s *userAccountLoginService) Execute(ctx context.Context, param *data.LoginUserAccountParam) (*data.LoginUserAccountResult, error) {
	userAccount, err := s.repo.FindByUsername(ctx, param.UsernameOrEmail)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userAccount.PasswordDigest), []byte(param.Password)); err != nil {
		return nil, err
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":     now.Unix(),
		"exp":     now.Add(time.Hour * 24).Unix(),
		"user_id": fmt.Sprintf("%d", userAccount.ID),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	result := data.NewLoginUserAccountResult(userAccount, tokenString)
	return result, nil
}
