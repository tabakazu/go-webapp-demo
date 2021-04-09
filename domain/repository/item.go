package repository

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
)

type Item interface {
	FindAll() (domain.ItemList, error)
	Find(domain.ItemID) (domain.Item, error)
	Create(*domain.Item) error
	UpdateAttributes(*domain.Item, map[string]interface{}) error
	SoftDelete(*domain.Item) error
}
