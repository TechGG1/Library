package service

import (
	"context"
	"github.com/TechGG1/Library/internal/model"
	"math"
	"time"
)

func (s *Service) CalculateFine(ctx context.Context, rentId int) (*model.Rent, error) {
	rentFromDB, err := s.LibraryRepo.RentById(ctx, rentId)
	rentFromDB.RentId = rentId
	if err != nil {
		return nil, err
	}
	if time.Now().Before(rentFromDB.LastDate) {
		rentFromDB.Fine = 0
		return rentFromDB, nil
	}
	pricePerDay, err := s.LibraryRepo.BookPricePerDay(ctx, rentFromDB.BookId)
	fineDays := time.Now().Sub(rentFromDB.LastDate).Hours() / 24
	rentFromDB.Fine = math.Round(fineDays) * float64(pricePerDay)

	return rentFromDB, nil
}
