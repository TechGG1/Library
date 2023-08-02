package service

import (
	"github.com/TechGG1/Library/internal/logging"
	"github.com/TechGG1/Library/internal/model"
)

//go:generate mockgen -source=service.go -destination=moks/mock.go

type Library interface {
	Books(limit, page int) ([]model.Book, error)
	CreateBook(book *model.Book) (int, error)
}

type LibraryRepo interface {
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
