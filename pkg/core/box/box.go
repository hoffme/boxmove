package box

type Data struct {
	Name  string
	Route []string
	Meta  interface{}
}

type Box interface {
	ID() string
	Data() (*Data, error)
}
