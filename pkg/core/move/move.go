package move

import "time"

type Data struct {
	FromID   string
	ToID     string
	Date     time.Time
	ItemID   string
	Quantity uint64
	Meta     interface{}
}

type Move interface {
	ID() string
	Data() (*Data, error)
}
