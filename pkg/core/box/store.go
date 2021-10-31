package box

type StoreSearchParams struct {
	ID         []string
	Name       string
	ParentID   string
	AncestorID string
}

type StoreCreateParams struct {
	Name   string
	Parent string
	Meta   interface{}
}

type Store interface {
	Find(id string) (Box, error)
	Search(params *StoreSearchParams) ([]Box, error)
	Create(params *StoreCreateParams) (Box, error)
}
