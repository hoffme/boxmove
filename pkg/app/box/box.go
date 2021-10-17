package box

type Data struct {
	ID       string
	Name     string
	ParentID string
}

type Filter struct {
	IDs      []string
	Name     string
	ParentID string
}

type CreateParams struct {
	Name   string
	Parent string
}

type Box interface {
	Data() (*Data, error)
}
