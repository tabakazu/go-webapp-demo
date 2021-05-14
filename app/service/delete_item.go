package service

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type deleteItem struct {
	Repo domain.ItemRepository
}

func NewDeleteItem(r domain.ItemRepository) *deleteItem {
	return &deleteItem{r}
}

func (s *deleteItem) Execute(itemId value.ItemID) (entity.Item, error) {
	item, err := s.Repo.Find(itemId)
	if err != nil {
		return item, err
	}

	if err := s.Repo.Delete(&item); err != nil {
		return item, err
	}

	return item, nil
}
