package application

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/repository"
)

type itemListService struct {
	Repository repository.Item
}

func NewItemListService(r repository.Item) itemListService {
	return itemListService{r}
}

func (s itemListService) Execute() ([]domain.Item, error) {
	items, err := s.Repository.FindAll()
	if err != nil {
		return items, err
	}

	return items, nil
}
