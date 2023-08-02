package model

import "time"

type Book struct {
	BookId      int       `json:"book_id"`
	Name        string    `json:"name"`
	Genre       []Genre   `json:"genre"`
	PriceOfBook int       `json:"price_of_book"`
	NumOfCopies int       `json:"num_of_copies"`
	Authors     string    `json:"authors"`
	CoverPhoto  string    `json:"cover_photo"`
	PricePerDay int       `json:"price_per_day"`
	RegDate     time.Time `json:"reg_date"`
}
