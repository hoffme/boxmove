package item

type StoreSearchParams struct {
	ID   []string
	Name string
}

type StoreCreateParams struct {
	Name string
	Meta interface{}
}

type Store interface {
	Find(id string) (Item, error)
	Search(params *StoreSearchParams) ([]Item, error)
	Create(params *StoreCreateParams) (Item, error)
}
