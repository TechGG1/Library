package service

import (
	"context"
	"errors"
	"github.com/TechGG1/Library/internal/logging"
	"github.com/TechGG1/Library/internal/model"
	mockService "github.com/TechGG1/Library/internal/service/moks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_Books(t *testing.T) {
	type mockBehaviorRepo func(s *mockService.MockLibraryRepo, limit, page int)
	testTable := []struct {
		name             string
		page             int
		limit            int
		mockBehaviorRepo mockBehaviorRepo
		expectedError    error
		expectedPage     int
		expectedBooks    []model.Book
	}{
		{
			name:  "OK",
			limit: 10,
			page:  1,
			mockBehaviorRepo: func(s *mockService.MockLibraryRepo, limit, page int) {
				s.EXPECT().BooksWithPage(gomock.Any(), limit, page).Return([]model.Book{
					{
						BookId: 1,
						Name:   "book1",
						Genre: []model.Genre{
							{
								Id:   1,
								Name: "horror",
							},
							{
								Id:   2,
								Name: "thriller",
							},
						},
						PriceOfBook: 100,
						PricePerDay: 15,
						NumOfCopies: 2,
						Authors:     "author1, author2",
						CoverPhoto:  "empty",
						RegDate:     time.Time{},
					},
					{
						BookId: 2,
						Name:   "book2",
						Genre: []model.Genre{
							{
								Id:   1,
								Name: "horror",
							},
						},
						PriceOfBook: 200,
						PricePerDay: 30,
						NumOfCopies: 5,
						Authors:     "author1",
						CoverPhoto:  "empty",
						RegDate:     time.Time{},
					},
				}, 1, nil)
			},
			expectedError: nil,
			expectedPage:  1,
			expectedBooks: []model.Book{
				{
					BookId: 1,
					Name:   "book1",
					Genre: []model.Genre{
						{
							Id:   1,
							Name: "horror",
						},
						{
							Id:   2,
							Name: "thriller",
						},
					},
					PriceOfBook: 100,
					PricePerDay: 15,
					NumOfCopies: 2,
					Authors:     "author1, author2",
					CoverPhoto:  "empty",
					RegDate:     time.Time{},
				},
				{
					BookId: 2,
					Name:   "book2",
					Genre: []model.Genre{
						{
							Id:   1,
							Name: "horror",
						},
					},
					PriceOfBook: 200,
					PricePerDay: 30,
					NumOfCopies: 5,
					Authors:     "author1",
					CoverPhoto:  "empty",
					RegDate:     time.Time{},
				},
			},
		},
		{
			name:  "Error in db",
			limit: 10,
			page:  1,
			mockBehaviorRepo: func(s *mockService.MockLibraryRepo, limit, page int) {
				s.EXPECT().BooksWithPage(gomock.Any(), limit, page).Return(nil, -1, errors.New("some error in db"))
			},
			expectedPage:  -1,
			expectedError: errors.New("some error in db"),
			expectedBooks: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockService.NewMockLibraryRepo(c)
			logger := logging.NewMockLogger()

			testCase.mockBehaviorRepo(repo, testCase.limit, testCase.page)

			serv := NewService(repo, logger)
			ctx := context.Background()

			books, _, err := serv.Books(ctx, testCase.limit, testCase.page)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedBooks, books)

			ctx.Done()
		})
	}
}

func TestService_CreateBook(t *testing.T) {
	type mockBehavior func(s *mockService.MockLibraryRepo, book *model.Book)
	testTable := []struct {
		name          string
		mockBook      *model.Book
		mockBehavior  mockBehavior
		expectedError error
		expectedId    int
	}{
		{
			name: "OK",
			mockBook: &model.Book{
				Name: "book1",
				Genre: []model.Genre{
					{
						Id:   1,
						Name: "horror",
					},
					{
						Id:   2,
						Name: "thriller",
					},
				},
				PriceOfBook: 100,
				PricePerDay: 15,
				NumOfCopies: 2,
				Authors:     "author1, author2",
				CoverPhoto:  "empty",
				RegDate:     time.Time{},
			},
			mockBehavior: func(s *mockService.MockLibraryRepo, book *model.Book) {
				s.EXPECT().CreateBook(gomock.Any(), book).Return(1, nil)
			},
			expectedError: nil,
			expectedId:    1,
		},
		{
			name: "Empty fields",
			mockBook: &model.Book{
				PriceOfBook: 100,
				PricePerDay: 15,
			},
			mockBehavior:  func(s *mockService.MockLibraryRepo, book *model.Book) {},
			expectedError: errors.New("missed information, empty fields"),
			expectedId:    -1,
		},
		{
			name: "Error in db",
			mockBook: &model.Book{
				Name: "book1",
				Genre: []model.Genre{
					{
						Id:   1,
						Name: "horror",
					},
				},
				PriceOfBook: 100,
				PricePerDay: 15,
				NumOfCopies: 2,
				Authors:     "author1, author2",
				CoverPhoto:  "empty",
				RegDate:     time.Time{},
			},
			mockBehavior: func(s *mockService.MockLibraryRepo, book *model.Book) {
				s.EXPECT().CreateBook(gomock.Any(), book).Return(0, errors.New("some error in db"))
			},
			expectedId:    -1,
			expectedError: errors.New("some error in db"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockService.NewMockLibraryRepo(c)

			testCase.mockBehavior(repo, testCase.mockBook)
			logger := logging.NewMockLogger()

			serv := NewService(repo, logger)
			ctx := context.Background()
			id, err := serv.CreateBook(ctx, testCase.mockBook)

			//Assert
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedId, id)
		})
	}
}
