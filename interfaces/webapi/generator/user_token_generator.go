package generator

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tabakazu/go-webapp/domain"
	"github.com/tabakazu/go-webapp/domain/entity"
)

type userTokenGenerator struct{}

func NewUserTokenGenerator() domain.UserTokenGenerator {
	return &userTokenGenerator{}
}

func (g userTokenGenerator) Issue(e *entity.UserAccount) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":     now.Unix(),
		"exp":     now.Add(time.Hour * 24).Unix(),
		"user_id": fmt.Sprintf("%d", e.ID),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
