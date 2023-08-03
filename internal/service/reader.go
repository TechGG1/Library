package service

import (
	"context"
	"github.com/TechGG1/Library/internal/model"
)

func (s *Service) CreateReader(ctx context.Context, reader *model.Reader) (int, error) {
	id, err := s.LibraryRepo.CreateReader(ctx, reader)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Service) Readers(ctx context.Context, limit, page int) ([]model.Reader, int, error) {
	readers, page, err := s.LibraryRepo.ReadersWithPage(ctx, limit, page)
	if err != nil {
		return nil, -1, err
	}
	return readers, page, nil
}

func (s *Service) UpdateReader(ctx context.Context, reader *model.Reader) (int, error) {
	id, err := s.LibraryRepo.UpdateReader(ctx, reader)
	if err != nil {
		return -1, err
	}
	return id, nil
}
