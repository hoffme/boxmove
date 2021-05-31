package box

import (
	"time"
)

type CreateParams struct {
	Name 	  string   	  	   `json:"name"`
	Meta 	  interface{} 	   `json:"meta"`
}

type UpdateParams struct {
	Route	  *[]string   	    `json:"route"`
	Name 	  *string   	    `json:"name"`
	Meta 	  *interface{} 	    `json:"meta"`
	Actives   *map[string]int64 `json:"actives"`
}

type Filter struct {
	ID         []string	`json:"id"`
	ParentID   *string  `json:"parent_id"`
	AncestorID string	`json:"ancestor_id"`
	Active     string   `json:"active"`
	Name       string	`json:"name"`
	Deleted    bool		`json:"deleted"`
}

type View struct {
	ID 		  string   	  	   `json:"id"`
	Route	  []string 	  	   `json:"route"`
	Name 	  string   	  	   `json:"name"`
	Meta 	  interface{} 	   `json:"meta"`
	Actives   map[string]int64 `json:"actives"`
	CreatedAt time.Time   	   `json:"created_at"`
	UpdatedAt time.Time   	   `json:"updated_at"`
	DeletedAt *time.Time  	   `json:"deleted_at,omitempty"`
}
