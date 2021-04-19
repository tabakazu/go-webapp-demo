package entity

import (
	"time"

	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type User struct {
	ID                   uint
	Email                value.Email
	Password             value.Password       `gorm:"-" json:"-"`
	PasswordConfirmation value.Password       `gorm:"-" json:"-"`
	PasswordDigest       value.PasswordDigest `json:"-"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
