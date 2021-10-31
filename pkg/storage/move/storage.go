package move

import "time"

type DTO struct {
	ID       string
	Date     time.Time
	FromID   string
	ToID     string
	ItemID   string
	Quantity uint64
	Meta     interface{}
}

type Fields struct {
	FromID   string
	ToID     string
	Date     time.Time
	ItemID   string
	Quantity uint64
	Meta     interface{}
}

type Filter struct {
	ID          []string
	ItemID      string
	FromID      string
	ToID        string
	DateFrom    time.Time
	DateTo      time.Time
	QuantityMin uint64
	QuantityMax uint64

	Order     string
	Ascendent bool
	Start     uint
	Limit     uint
}

type Storage interface {
	Find(id string) (*DTO, error)
	Search(filter *Filter) ([]*DTO, error)
	Create(fields *Fields) (*DTO, error)
	Update(id string, fields *Fields) (*DTO, error)
	Delete(id string) error
}
