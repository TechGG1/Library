package model

type Rent struct {
	RentId    string  `json:"rent_id"`
	BookId    string  `json:"book_id"`
	ReaderId  string  `json:"reader_id"`
	FirstDate string  `json:"first_date"`
	LastDate  string  `json:"last_date"`
	Fine      float64 `json:"fine"`
}
