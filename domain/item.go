package domain

import "time"

type ItemID uint

type Item struct {
	ID        ItemID
	Name      string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
