package client

import (
	"time"
)

type CreateParams struct {
	Name   string      `json:"name"`
}

type UpdateParams struct {
	Name *string      `json:"name"`
}

type Filter struct {
	ID      []string `json:"id"`
	Name    string	 `json:"name"`
	Deleted bool	 `json:"deleted"`
}

type View struct {
	ID        string      `json:"id"`
	Name  	  string 	  `json:"name"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at,omitempty"`
}