package domain

type ItemID int

type Item struct {
	ID     ItemID
	Name   string
	Amount int
}
