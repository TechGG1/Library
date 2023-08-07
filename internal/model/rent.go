package model

import "time"

type Rent struct {
	RentId    int       `json:"id"`
	BookId    int       `json:"book_id"`
	ReaderId  int       `json:"reader_id"`
	FirstDate time.Time `json:"first_date"`
	LastDate  time.Time `json:"last_date"`
	Fine      float64   `json:"fine"`
	Complete  bool      `json:"complete"`
}
