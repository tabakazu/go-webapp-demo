package interactor

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type itemUpdate struct {
	Repository domain.ItemRepository
}

func NewItemUpdate(r domain.ItemRepository) itemUpdate {
	return itemUpdate{r}
}

func (s itemUpdate) Execute(itemId value.ItemID, newItem entity.Item) (entity.Item, error) {
	item, err := s.Repository.Find(itemId)
	if err != nil {
		return item, err
	}

	attributes := map[string]interface{}{
		"Name":   newItem.Name,
		"Amount": newItem.Amount,
	}
	if err := s.Repository.UpdateAttributes(&item, attributes); err != nil {
		return item, err
	}
	return item, nil
}
