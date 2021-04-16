package interactor

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type itemGet struct {
	Repository domain.ItemRepository
}

func NewItemGet(r domain.ItemRepository) itemGet {
	return itemGet{r}
}

func (s itemGet) Execute(itemId value.ItemID) (entity.Item, error) {
	item, err := s.Repository.Find(itemId)
	if err != nil {
		return item, err
	}

	return item, nil
}
