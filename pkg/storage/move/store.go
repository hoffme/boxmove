package move

import "time"

type DTO struct {
	ID       string
	FromID   string
	ToID     string
	Date     time.Time
	ItemID   string
	Quantity uint64
}

type Fields struct {
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

type UpdateMany struct {
	Id     string
	Fields *Fields
}

type Store interface {
	Find(id string) (*DTO, error)
	Search(filter *Filter) ([]*DTO, error)
	Create(fields *Fields) (*DTO, error)
	Update(id string, fields *Fields) (*DTO, error)
	Delete(id string) error
}
