package move

import "time"

type StoreSearchParams struct {
	ID          []string
	ItemID      string
	FromID      string
	ToID        string
	DateFrom    time.Time
	DateTo      time.Time
	QuantityMin uint64
	QuantityMax uint64
}

type StoreCreateParams struct {
	FromID   string
	ToID     string
	ItemID   string
	Quantity uint64
	Meta     interface{}
}

type Store interface {
	Find(id string) (Move, error)
	Search(params *StoreSearchParams) ([]Move, error)
	Create(params *StoreCreateParams) (Move, error)
}
