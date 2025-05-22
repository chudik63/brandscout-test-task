package service

import (
	"brandscout-test-task/internal/models"
	"brandscout-test-task/internal/repository"
	"brandscout-test-task/internal/service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllQuotes(t *testing.T) {
	tests := []struct {
		name           string
		authorParam    string
		expectedStatus int
		mockRepo       func(m *mocks.QuotesRepository)
		expectedQuotes []*models.Quote
		expectedError  error
	}{
		{
			name:        "success",
			authorParam: "",
			mockRepo: func(m *mocks.QuotesRepository) {
				m.On("GetAllQuotes", context.Background()).Return([]*models.Quote{
					{ID: 1, Author: "Author1", Quote: "Quote1"},
					{ID: 2, Author: "Author2", Quote: "Quote2"},
				}, nil)
			},
			expectedQuotes: []*models.Quote{
				{ID: 1, Author: "Author1", Quote: "Quote1"},
				{ID: 2, Author: "Author2", Quote: "Quote2"},
			},
			expectedError: nil,
		},
		{
			name:        "Author1",
			authorParam: "Author1",
			mockRepo: func(m *mocks.QuotesRepository) {
				m.On("GetQuotesByAuthor", context.Background(), "Author1").Return([]*models.Quote{
					{ID: 1, Author: "Author1", Quote: "Quote1"},
				}, nil)
			},
			expectedQuotes: []*models.Quote{
				{ID: 1, Author: "Author1", Quote: "Quote1"},
			},
			expectedError: nil,
		},
		{
			name:        "Not found",
			authorParam: "Author3",
			mockRepo: func(m *mocks.QuotesRepository) {
				m.On("GetQuotesByAuthor", context.Background(), "Author3").Return(nil, repository.ErrNotFound)
			},
			expectedQuotes: nil,
			expectedError:  repository.ErrNotFound,
		},
		{
			name:        "Not found",
			authorParam: "",
			mockRepo: func(m *mocks.QuotesRepository) {
				m.On("GetAllQuotes", context.Background()).Return(nil, repository.ErrNotFound)
			},
			expectedQuotes: nil,
			expectedError:  repository.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.QuotesRepository)
			s := New(mockRepo)

			tt.mockRepo(mockRepo)

			res, err := s.GetAllQuotes(context.Background(), tt.authorParam)

			assert.Equal(t, tt.expectedQuotes, res)
			assert.Equal(t, tt.expectedError, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
