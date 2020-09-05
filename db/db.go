package db

import (
	"context"

	"github.com/tabakazu/golang-webapi-demo/model"
)

type DB interface {
	GetAllItems(ctx context.Context) ([]*model.Item, error)
	PutItem(ctx context.Context, t *model.Item) error
}
