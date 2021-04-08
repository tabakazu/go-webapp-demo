package application

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/repository"
)

type itemShowService struct {
	Repository repository.Item
}

func NewItemShowService(r repository.Item) itemShowService {
	return itemShowService{r}
}

func (s itemShowService) Execute(itemId domain.ItemID) (domain.Item, error) {
	item, err := s.Repository.Find(itemId)
	if err != nil {
		return item, err
	}

	return item, nil
}
