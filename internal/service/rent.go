package service

import (
	"context"
	"github.com/TechGG1/Library/internal/model"
)

func (s *Service) CreateRent(ctx context.Context, rent *model.Rent) (int, error) {
	id, err := s.LibraryRepo.CreateRent(ctx, rent)
	if err != nil {
		return -1, err
	}
	return id, nil
}
