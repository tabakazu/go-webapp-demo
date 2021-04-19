package gateway

import (
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(d *gorm.DB) UserRepository {
	return UserRepository{d}
}

func (r UserRepository) FindByEmail(email value.Email) (entity.User, error) {
	var user entity.User
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) Create(user *entity.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
