package repository

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
)

type Item interface {
	Find(domain.ItemID) (domain.Item, error)
}
