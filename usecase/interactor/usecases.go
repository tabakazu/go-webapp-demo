package interactor

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/usecase"
)

type ItemsUseCases struct {
	GetList usecase.ItemsGet
	Get     usecase.ItemGet
	Create  usecase.ItemCreate
	Update  usecase.ItemUpdate
	Delete  usecase.ItemDelete
}

func NewItemsUseCases(r domain.ItemRepository) ItemsUseCases {
	return ItemsUseCases{
		GetList: NewItemsGet(r),
		Get:     NewItemGet(r),
		Create:  NewItemCreate(r),
		Update:  NewItemUpdate(r),
		Delete:  NewItemDelete(r),
	}
}
