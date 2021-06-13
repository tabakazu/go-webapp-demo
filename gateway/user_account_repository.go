package gateway

import (
	"context"

	"github.com/tabakazu/go-webapp/domain"
	"gorm.io/gorm"
)

type userAccountRepository struct {
	db *gorm.DB
}

func NewUserAccountRepository(db *gorm.DB) domain.UserAccountRepository {
	return &userAccountRepository{db: db}
}

func (r *userAccountRepository) Create(ctx context.Context, e *domain.UserAccount) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	u := domain.User{
		Username:   e.Username,
		FamilyName: e.FamilyName,
		GivenName:  e.GivenName,
	}
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	e.ID = u.ID

	a := domain.Account{
		UserID:         e.ID,
		Email:          e.Email,
		PasswordDigest: e.PasswordDigest,
	}
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
