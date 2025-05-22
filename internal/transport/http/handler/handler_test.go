package handler

import (
	"brandscout-test-task/internal/models"
	"brandscout-test-task/internal/repository"
	"brandscout-test-task/internal/transport/http/handler/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetQuotes(t *testing.T) {
	tests := []struct {
		name           string
		authorParam    string
		expectedStatus int
		mockQuotes     []*models.Quote
		mockError      error
	}{
		{
			name:           "success",
			authorParam:    "Confucius",
			expectedStatus: http.StatusOK,
			mockQuotes: []*models.Quote{
				{Author: "Confucius", Quote: "Wisdom"},
			},
			mockError: nil,
		},
		{
			name:           "success",
			authorParam:    "Confucius",
			expectedStatus: http.StatusNotFound,
			mockQuotes:     nil,
			mockError:      repository.ErrNotFound,
		},
		{
			name:           "internal error",
			authorParam:    "",
			expectedStatus: http.StatusInternalServerError,
			mockQuotes:     nil,
			mockError:      errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := new(mocks.Service)
			mockSvc.On("GetAllQuotes", mock.Anything, tt.authorParam).Return(tt.mockQuotes, tt.mockError)

			h := New(mockSvc)

			req := httptest.NewRequest(http.MethodGet, "/quotes?author="+tt.authorParam, nil)
			rec := httptest.NewRecorder()

			h.GetQuotes(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			mockSvc.AssertExpectations(t)
		})
	}
}

func TestDeleteQuote(t *testing.T) {
	tests := []struct {
		name           string
		quoteID        string
		mockService    func(m *mocks.Service)
		expectedStatus int
	}{
		{
			name:    "success",
			quoteID: "42",
			mockService: func(m *mocks.Service) {
				m.On("DeleteQuote", mock.Anything, uint64(42)).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:    "invalid id",
			quoteID: "invalid_number",
			mockService: func(m *mocks.Service) {

			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:    "not found",
			quoteID: "99",
			mockService: func(m *mocks.Service) {
				m.On("DeleteQuote", mock.Anything, uint64(99)).Return(repository.ErrNotFound)
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:    "internal error",
			quoteID: "100",
			mockService: func(m *mocks.Service) {
				m.On("DeleteQuote", mock.Anything, uint64(100)).Return(errors.New("db error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := new(mocks.Service)
			tt.mockService(mockSvc)

			h := New(mockSvc)

			req := httptest.NewRequest(http.MethodGet, "/quotes/"+tt.quoteID, nil)
			rec := httptest.NewRecorder()

			h.DeleteQuote(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			mockSvc.AssertExpectations(t)
		})
	}
}
