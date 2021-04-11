package gateway

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(d *gorm.DB) ItemRepository {
	return ItemRepository{d}
}

func (r ItemRepository) FindAll() (domain.ItemList, error) {
	var items domain.ItemList
	if err := r.DB.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r ItemRepository) Find(itemId domain.ItemID) (domain.Item, error) {
	var item domain.Item
	if err := r.DB.First(&item, "id = ?", itemId).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r ItemRepository) Create(item *domain.Item) error {
	if err := r.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (r ItemRepository) UpdateAttributes(item *domain.Item, attributes map[string]interface{}) error {
	if err := r.DB.Model(item).Updates(attributes).Error; err != nil {
		return err
	}
	return nil
}

func (r ItemRepository) Delete(item *domain.Item) error {
	if err := r.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}
