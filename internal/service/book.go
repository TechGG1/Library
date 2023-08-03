package service

import (
	"context"
	"github.com/TechGG1/Library/internal/model"
	"time"
)

func (s *Service) Books(ctx context.Context, limit, page int) ([]model.Book, int, error) {
	books, page, err := s.LibraryRepo.BooksWithPage(ctx, limit, page)
	if err != nil {
		return nil, -1, err
	}
	return books, page, nil
}

func (s *Service) CreateBook(ctx context.Context, book *model.Book) (int, error) {
	book.RegDate = time.Now()
	id, err := s.LibraryRepo.CreateBook(ctx, book)
	if err != nil {
		return -1, err
	}
	return id, nil
}
