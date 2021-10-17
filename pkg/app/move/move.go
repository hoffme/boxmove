package move

import "time"

type Data struct {
	ID       string
	FromID   string
	ToID     string
	Date     time.Time
	ItemID   string
	Quantity uint64
}

type Filter struct {
	IDs      []string
	ItemID   string
	FromID   string
	ToID     string
	FromDate time.Time
	ToDate   time.Time
	Min      uint64
	Max      uint64
}

type CreateParams struct {
	FromID   string
	ToID     string
	ItemID   string
	Quantity uint64
}

type Move interface {
	Data() (*Data, error)
}
