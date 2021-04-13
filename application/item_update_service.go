package application

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/domain/repository"
)

type itemUpdateService struct {
	Repository repository.Item
}

func NewItemUpdateService(r repository.Item) itemUpdateService {
	return itemUpdateService{r}
}

func (s itemUpdateService) Execute(itemId domain.ItemID, newItem domain.Item) (domain.Item, error) {
	item, err := s.Repository.Find(itemId)
	if err != nil {
		return item, err
	}

	attributes := map[string]interface{}{
		"Name":   newItem.Name,
		"Amount": newItem.Amount,
	}
	if err := s.Repository.UpdateAttributes(&item, attributes); err != nil {
		return item, err
	}
	return item, nil
}
