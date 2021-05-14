package domain

import (
	"github.com/tabakazu/golang-webapi-demo/domain/collection"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type ItemRepository interface {
	FindAll() (collection.Items, error)
	Find(value.ItemID) (entity.Item, error)
	Create(*entity.Item) error
	Delete(*entity.Item) error
}

type UserRepository interface {
	FindByEmail(value.Email) (entity.User, error)
	Create(*entity.User) error
}
