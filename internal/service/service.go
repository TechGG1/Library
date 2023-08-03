package service

import (
	"context"
	"github.com/TechGG1/Library/internal/logging"
	"github.com/TechGG1/Library/internal/model"
)

//go:generate mockgen -source=service.go -destination=moks/mock.go

type Library interface {
	Books(ctx context.Context, limit, page int) ([]model.Book, int, error)
	CreateBook(ctx context.Context, book *model.Book) (int, error)
}

type LibraryRepo interface {
	BooksWithPage(ctx context.Context, limit, page int) ([]model.Book, int, error)
	CreateBook(ctx context.Context, book *model.Book) (int, error)
}

type Service struct {
	LibraryRepo
	Log *logging.Logger
}

func NewService(rep LibraryRepo, log *logging.Logger) *Service {
	return &Service{
		LibraryRepo: rep,
		Log:         log,
	}
}
