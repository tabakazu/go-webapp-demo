package db

import (
	"context"

	"github.com/tabakazu/golang-webapi-demo/model"
)

type DB interface {
	GetAllItems(ctx context.Context) ([]*model.Item, error)
	CreateItem(ctx context.Context, t *model.Item) error
}
