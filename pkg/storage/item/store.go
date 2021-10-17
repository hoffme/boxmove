package item

type DTO struct {
	ID   string
	Name string
}

type Fields struct {
	Name string
}

type Filter struct {
	IDs  []string
	Name string
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
