package usecase

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
)

type ItemShow interface {
	Execute(domain.ItemID) (domain.Item, error)
}

type ItemList interface {
	Execute() ([]domain.Item, error)
}

type ItemCreate interface {
	Execute(domain.Item) (domain.Item, error)
}

type ItemUpdate interface {
	Execute(domain.ItemID, domain.Item) (domain.Item, error)
}

type ItemDelete interface {
	Execute(domain.ItemID) (domain.Item, error)
}
