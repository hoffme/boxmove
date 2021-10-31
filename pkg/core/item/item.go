package item

type Data struct {
	Name string
	Meta interface{}
}

type Item interface {
	ID() string
	Data() (*Data, error)
}
