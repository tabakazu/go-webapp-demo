package usecase

import (
	"github.com/tabakazu/golang-webapi-demo/domain/collection"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type ItemsGet interface {
	Execute() (collection.Items, error)
}

type ItemGet interface {
	Execute(value.ItemID) (entity.Item, error)
}

type ItemCreate interface {
	Execute(entity.Item) (entity.Item, error)
}

type ItemUpdate interface {
	Execute(value.ItemID, entity.Item) (entity.Item, error)
}

type ItemDelete interface {
	Execute(value.ItemID) (entity.Item, error)
}
