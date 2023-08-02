package model

type Reader struct {
	IdReader    int    `json:"reader_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	Email       string `json:"email"`
}
