package interactor

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
)

type itemCreate struct {
	Repository domain.ItemRepository
}

func NewItemCreate(r domain.ItemRepository) itemCreate {
	return itemCreate{r}
}

func (s itemCreate) Execute(item entity.Item) (entity.Item, error) {
	if err := s.Repository.Create(&item); err != nil {
		return item, err
	}

	return item, nil
}
