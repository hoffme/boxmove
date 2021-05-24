package move

import "time"

type Filter struct {
	ID       []string  `json:"id"`
	FromID   []string  `json:"from_id"`
	ToID     []string  `json:"to_id"`
	CountMin uint64    `json:"count_min"`
	CountMax uint64    `json:"count_max"`
	DateMin  time.Time `json:"date_min"`
	DateMax  time.Time `json:"date_max"`
}

type CreateParams struct {
	FromID string 	 `json:"from_id"`
	ToID   string 	 `json:"to_id"`
	Date   time.Time `json:"date"`
	Count  uint64 	 `json:"count"`
}