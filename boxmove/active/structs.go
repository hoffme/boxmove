package active

import (
	"time"
)

type CreateParams struct {
	Code   string 	   `json:"code"`
	Name   string      `json:"name"`
	Meta   interface{} `json:"meta"`
}

type UpdateParams struct {
	Name *string      `json:"name"`
	Meta *interface{} `json:"meta"`
}

type Filter struct {
	ID      []string `json:"id"`
	Codes  	[]string `json:"codes"`
	Name    string	 `json:"name"`
	Deleted bool	 `json:"deleted"`
}

type View struct {
	ID        string      `json:"id"`
	Code      string      `json:"code"`
	Name  	  string 	  `json:"name"`
	Meta 	  interface{} `json:"meta"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
}