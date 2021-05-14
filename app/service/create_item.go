package service

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
)

type createItem struct {
	Repo domain.ItemRepository
}

func NewCreateItem(r domain.ItemRepository) *createItem {
	return &createItem{r}
}

func (s *createItem) Execute(item entity.Item) (entity.Item, error) {
	if err := s.Repo.Create(&item); err != nil {
		return item, err
	}

	return item, nil
}
