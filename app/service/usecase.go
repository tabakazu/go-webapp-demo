package service

import (
	"github.com/tabakazu/golang-webapi-demo/app"
	"github.com/tabakazu/golang-webapi-demo/domain"
)

type ItemUseCase struct {
	GetList app.GetItemListUseCase
	Get     app.GetItemUseCase
	Create  app.CreateItemUseCase
	Delete  app.DeleteItemUseCase
}

func NewItemUseCase(r domain.ItemRepository) *ItemUseCase {
	return &ItemUseCase{
		GetList: NewGetItemList(r),
		Get:     NewGetItem(r),
		Create:  NewCreateItem(r),
		Delete:  NewDeleteItem(r),
	}
}
