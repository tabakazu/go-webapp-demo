package app

import (
	"github.com/tabakazu/golang-webapi-demo/domain/collection"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type GetItemListUseCase interface {
	Execute() (collection.Items, error)
}

type GetItemUseCase interface {
	Execute(value.ItemID) (entity.Item, error)
}

type CreateItemUseCase interface {
	Execute(entity.Item) (entity.Item, error)
}

type DeleteItemUseCase interface {
	Execute(value.ItemID) (entity.Item, error)
}
