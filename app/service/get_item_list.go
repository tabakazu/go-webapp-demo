package service

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/collection"
)

type getItemList struct {
	Repo domain.ItemRepository
}

func NewGetItemList(r domain.ItemRepository) *getItemList {
	return &getItemList{r}
}

func (s getItemList) Execute() (collection.Items, error) {
	items, err := s.Repo.FindAll()
	if err != nil {
		return items, err
	}

	return items, nil
}
