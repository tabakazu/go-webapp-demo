package interactor

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type itemDelete struct {
	Repository domain.ItemRepository
}

func NewItemDelete(r domain.ItemRepository) itemDelete {
	return itemDelete{r}
}

func (s itemDelete) Execute(itemId value.ItemID) (entity.Item, error) {
	item, err := s.Repository.Find(itemId)
	if err != nil {
		return item, err
	}

	if err := s.Repository.Delete(&item); err != nil {
		return item, err
	}

	return item, nil
}
