package interactor

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/collection"
)

type itemsGet struct {
	Repository domain.ItemRepository
}

func NewItemsGet(r domain.ItemRepository) itemsGet {
	return itemsGet{r}
}

func (s itemsGet) Execute() (collection.Items, error) {
	items, err := s.Repository.FindAll()
	if err != nil {
		return items, err
	}

	return items, nil
}
