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

func (r *userAccountRepository) userAccountScope() *gorm.DB {
	return r.db.Model(&domain.User{}).
		Select("users.username, users.family_name, users.given_name, a.email, a.password_digest").
		Joins("JOIN accounts a ON a.user_id = users.id")
}

func (r *userAccountRepository) FindByUsername(ctx context.Context, username string) (*domain.UserAccount, error) {
	var e domain.UserAccount
	if err := r.userAccountScope().Where("username = ?", username).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *userAccountRepository) FindByEmail(ctx context.Context, email string) (*domain.UserAccount, error) {
	var e domain.UserAccount
	if err := r.userAccountScope().Where("email = ?", email).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
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
