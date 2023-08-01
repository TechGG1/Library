package service

import (
	"library/internal/logging"
)

//go:generate mockgen -source=service.go -destination=moks/mock.go

type Library interface {
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
