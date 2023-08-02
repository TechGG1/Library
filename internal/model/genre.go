package model

type Genre struct {
	Id   int    `json:"genre_id"`
	Name string `json:"name"`
}

type GenreToBook struct {
	GenreId int `json:"genre_id"`
	BookId  int `json:"book_id"`
}
