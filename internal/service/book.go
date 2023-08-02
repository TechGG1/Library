package service

import (
	"fmt"
	"github.com/TechGG1/Library/internal/model"
	"time"
)

func (s *Service) Books(limit, page int) ([]model.Book, error) {
	return nil, nil
}

func (s *Service) CreateBook(book *model.Book) (int, error) {
	book.RegDate = time.Now()
	fmt.Println(book)
	return 1, nil
}
