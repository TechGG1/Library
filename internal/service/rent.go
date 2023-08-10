package service

import (
	"context"
	"errors"
	"github.com/TechGG1/Library/internal/model"
	"time"
)

func (s *Service) CreateRent(ctx context.Context, rent *model.Rent) (int, error) {
	rent.FirstDate = time.Now()
	rent.LastDate = time.Now().Add(time.Hour * 24 * 30)
	id, err := s.LibraryRepo.CreateRent(ctx, rent)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Service) UpdateRent(ctx context.Context, rent *model.Rent) (int, error) {
	rentWithFine, err := s.CalculateFine(ctx, rent.RentId)
	if err != nil {
		return -1, nil
	}
	id, err := s.LibraryRepo.UpdateRent(ctx, rentWithFine)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Service) Rents(ctx context.Context, limit, page, readerId int) ([]model.Rent, int, error) {
	if limit <= 0 || page <= 0 || readerId <= 0 {
		return nil, -1, errors.New("enter correct limit/page/reader_id")
	}
	rents, page, err := s.LibraryRepo.RentsWithPage(ctx, limit, page, readerId)
	if err != nil {
		return nil, -1, err
	}
	return rents, page, nil
}
