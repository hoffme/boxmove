package boxmove

import (
	"time"
)

type BoxMovesOptions struct {
	Ingress  bool		`json:"ingress"`
	Egress   bool		`json:"egress"`
	CountMin uint64		`json:"count_min"`
	CountMax uint64		`json:"count_max"`
	DateMin  time.Time	`json:"date_min"`
	DateMax  time.Time	`json:"date_max"`
	Deleted  bool		`json:"deleted"`
}

type BoxCountsOptions struct {
	DateMin  time.Time	`json:"date_min"`
	DateMax  time.Time	`json:"date_max"`
	Deleted  bool		`json:"deleted"`
}