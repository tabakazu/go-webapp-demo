package entity

import (
	"time"

	"github.com/tabakazu/go-webapp/domain/value"
)

type UserAccount struct {
	ID             uint
	Username       string
	FamilyName     string
	GivenName      string
	Email          string
	PasswordDigest value.PasswordDigest
}

func (e UserAccount) ValidPassword(passwd value.Password) error {
	return e.PasswordDigest.ValidPassword(passwd)
}

type User struct {
	ID         uint
	Username   string
	FamilyName string `gorm:"column:family_name"`
	GivenName  string `gorm:"column:given_name"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Account struct {
	ID             uint
	UserID         uint `gorm:"column:user_id"`
	Email          string
	PasswordDigest value.PasswordDigest `gorm:"column:password_digest"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
