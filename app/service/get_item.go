package service

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type getItem struct {
	Repo domain.ItemRepository
}

func NewGetItem(r domain.ItemRepository) *getItem {
	return &getItem{r}
}

func (s *getItem) Execute(itemId value.ItemID) (entity.Item, error) {
	item, err := s.Repo.Find(itemId)
	if err != nil {
		return item, err
	}

	return item, nil
}
