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
	CreateReader(ctx context.Context, reader *model.Reader) (int, error)
	Readers(ctx context.Context, limit, page int) ([]model.Reader, int, error)
	UpdateReader(ctx context.Context, reader *model.Reader) (int, error)
	CreateRent(ctx context.Context, rent *model.Rent) (int, error)
}

type LibraryRepo interface {
	BooksWithPage(ctx context.Context, limit, page int) ([]model.Book, int, error)
	CreateBook(ctx context.Context, book *model.Book) (int, error)
	CreateReader(ctx context.Context, reader *model.Reader) (int, error)
	ReadersWithPage(ctx context.Context, limit, page int) ([]model.Reader, int, error)
	UpdateReader(ctx context.Context, reader *model.Reader) (int, error)
	CreateRent(ctx context.Context, rent *model.Rent) (int, error)
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
