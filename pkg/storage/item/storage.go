package item

type DTO struct {
	ID   string
	Name string
	Meta interface{}
}

type Fields struct {
	Name string
	Meta interface{}
}

type Filter struct {
	ID   []string
	Name string

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
