package item

type QuantityBox struct {
	BoxID    string
	Ingress  uint64
	Egress   uint64
	Quantity uint64
}

type Data struct {
	ID   string
	Name string
}

type Filter struct {
	IDs         []string
	Name        string
	BoxID       string
	IngressMin  uint64
	IngressMax  uint64
	QuantityMin uint64
	QuantityMax uint64
	EgressMin   uint64
	EgressMax   uint64
}

type CreateParams struct {
	Name string
}

type Item interface {
	Data() (*Data, error)
	Quantity(boxID string) (*QuantityBox, error)
}
