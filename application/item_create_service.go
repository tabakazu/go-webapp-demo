package application

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/repository"
)

type itemCreateService struct {
	Repository repository.Item
}

func NewItemCreateService(r repository.Item) itemCreateService {
	return itemCreateService{r}
}

func (s itemCreateService) Execute(item domain.Item) (domain.Item, error) {
	if err := s.Repository.Create(&item); err != nil {
		return item, err
	}

	return item, nil
}
