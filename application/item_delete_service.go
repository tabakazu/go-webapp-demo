package application

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/repository"
)

type itemDeleteService struct {
	Repository repository.Item
}

func NewItemDeleteService(r repository.Item) itemDeleteService {
	return itemDeleteService{r}
}

func (s itemDeleteService) Execute(itemId domain.ItemID) (domain.Item, error) {
	item, err := s.Repository.Find(itemId)
	if err != nil {
		return item, err
	}

	if err := s.Repository.Delete(&item); err != nil {
		return item, err
	}

	return item, nil
}
