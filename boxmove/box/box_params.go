package box

type Filter struct {
	ID         []string	`json:"id"`
	ParentID   *string	`json:"parent_id"`
	AncestorID *string	`json:"ancestor_id"`
	Name       string	`json:"name"`
	Type       string	`json:"type"`
	Deleted    bool		`json:"deleted"`
}

type CreateParams struct {
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
