package repository

import (
	"context"

	"github.com/tabakazu/go-webapp/domain"
	"github.com/tabakazu/go-webapp/domain/entity"
	"gorm.io/gorm"
)

type userAccountRepository struct {
	db *gorm.DB
}

func NewUserAccountRepository(db *gorm.DB) domain.UserAccountRepository {
	return &userAccountRepository{db: db}
}

func userAccountScope(db *gorm.DB) *gorm.DB {
	return db.Model(&entity.User{}).
		Select("users.id, users.username, users.family_name, users.given_name, a.email, a.password_digest").
		Joins("JOIN accounts a ON a.user_id = users.id")
}

func (r *userAccountRepository) FindByID(ctx context.Context, userID int) (*entity.UserAccount, error) {
	var e entity.UserAccount
	if err := r.db.WithContext(ctx).
		Scopes(userAccountScope).
		Where("users.id = ?", userID).
		First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *userAccountRepository) FindByUsername(ctx context.Context, username string) (*entity.UserAccount, error) {
	var e entity.UserAccount
	if err := r.db.WithContext(ctx).
		Scopes(userAccountScope).
		Where("username = ?", username).
		First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *userAccountRepository) FindByEmail(ctx context.Context, email string) (*entity.UserAccount, error) {
	var e entity.UserAccount
	if err := r.db.WithContext(ctx).
		Scopes(userAccountScope).
		Where("email = ?", email).
		First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *userAccountRepository) Create(ctx context.Context, e *entity.UserAccount) error {
	tx := r.db.WithContext(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	u := entity.User{
		Username:   e.Username,
		FamilyName: e.FamilyName,
		GivenName:  e.GivenName,
	}
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	e.ID = u.ID

	a := entity.Account{
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
