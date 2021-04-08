package gateway

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(d *gorm.DB) ItemRepository {
	return ItemRepository{d}
}

func (r ItemRepository) Find(itemId domain.ItemID) (domain.Item, error) {
	var item domain.Item
	if err := r.db.First(&item, "id = ?", itemId).Error; err != nil {
		return item, err
	}
	return item, nil
}
