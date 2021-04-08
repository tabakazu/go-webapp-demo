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
