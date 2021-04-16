package entity

import (
	"time"

	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type Item struct {
	ID        value.ItemID
	Name      string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
