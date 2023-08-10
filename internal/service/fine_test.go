package service

import (
	"context"
	"github.com/TechGG1/Library/internal/logging"
	"github.com/TechGG1/Library/internal/model"
	mockService "github.com/TechGG1/Library/internal/service/moks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_Fine(t *testing.T) {
	type mockBehaviorRepoRent func(s *mockService.MockLibraryRepo, rentId int)
	type mockBehaviorRepoBook func(s *mockService.MockLibraryRepo, bookId int)
	testTable := []struct {
		name                 string
		mockRentId           int
		mockBookId           int
		mockBehaviorRepoRent mockBehaviorRepoRent
		mockBehaviorRepoBook mockBehaviorRepoBook
		expectedError        error
		expectedFine         float64
	}{
		{
			name:       "OK",
			mockRentId: 1,
			mockBookId: 1,
			mockBehaviorRepoRent: func(s *mockService.MockLibraryRepo, rentId int) {
				s.EXPECT().RentById(gomock.Any(), rentId).Return(&model.Rent{
					RentId:    1,
					BookId:    1,
					ReaderId:  1,
					FirstDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					LastDate:  time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local),
					Fine:      0,
					Complete:  false,
				}, nil)
			},
			mockBehaviorRepoBook: func(s *mockService.MockLibraryRepo, bookId int) {
				s.EXPECT().BookPricePerDay(gomock.Any(), bookId).Return(1, nil)
			},
			expectedFine:  192,
			expectedError: nil,
		},
		{
			name:       "Rent completed earlier",
			mockRentId: 1,
			mockBookId: 1,
			mockBehaviorRepoRent: func(s *mockService.MockLibraryRepo, rentId int) {
				s.EXPECT().RentById(gomock.Any(), rentId).Return(&model.Rent{
					RentId:    1,
					BookId:    1,
					ReaderId:  1,
					FirstDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
					LastDate:  time.Date(2023, 1, 31, 0, 0, 0, 0, time.Local),
					Fine:      0,
					Complete:  true,
				}, nil)
			},
			mockBehaviorRepoBook: func(s *mockService.MockLibraryRepo, bookId int) {},
			expectedFine:         0,
			expectedError:        nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockService.NewMockLibraryRepo(c)
			logger := logging.NewMockLogger()

			testCase.mockBehaviorRepoRent(repo, testCase.mockRentId)
			testCase.mockBehaviorRepoBook(repo, testCase.mockBookId)

			serv := NewService(repo, logger)
			ctx := context.Background()

			rent, err := serv.CalculateFine(ctx, testCase.mockRentId)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedFine, rent.Fine)

			ctx.Done()
		})
	}
}
