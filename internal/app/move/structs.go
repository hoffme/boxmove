package move

import (
	"time"
)

type CreateParams struct {
	Active string	   `json:"active"`
	From   string 	   `json:"from"`
	To     string	   `json:"to"`
	Date   time.Time   `json:"date"`
	Count  uint64  	   `json:"count"`
	Meta   interface{} `json:"meta"`
}

type Filter struct {
	ID       []string  `json:"id"`
	FromID   []string  `json:"from_id"`
	ToID     []string  `json:"to_id"`
	Active   string    `json:"active"`
	CountMin uint64    `json:"count_min"`
	CountMax uint64    `json:"count_max"`
	DateMin  time.Time `json:"date_min"`
	DateMax  time.Time `json:"date_max"`
}

type View struct {
	ID        string      `json:"id"`
	Active 	  string	  `json:"active"`
	From   	  string 	  `json:"from"`
	To     	  string	  `json:"to"`
	Date   	  time.Time   `json:"date"`
	Count  	  uint64  	  `json:"count"`
	Meta   	  interface{} `json:"meta"`
	CreatedAt time.Time   `json:"created_at"`
}