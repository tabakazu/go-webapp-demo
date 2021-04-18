package datastore

import (
	"github.com/tabakazu/golang-webapi-demo/domain/collection"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(d *gorm.DB) ItemRepository {
	return ItemRepository{d}
}

func (r ItemRepository) FindAll() (collection.Items, error) {
	var items collection.Items
	if err := r.DB.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

func (r ItemRepository) Find(itemId value.ItemID) (entity.Item, error) {
	var item entity.Item
	if err := r.DB.First(&item, "id = ?", itemId).Error; err != nil {
		return item, err
	}
	return item, nil
}

func (r ItemRepository) Create(item *entity.Item) error {
	if err := r.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (r ItemRepository) UpdateAttributes(item *entity.Item, attributes map[string]interface{}) error {
	if err := r.DB.Model(item).Updates(attributes).Error; err != nil {
		return err
	}
	return nil
}

func (r ItemRepository) Delete(item *entity.Item) error {
	if err := r.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}
