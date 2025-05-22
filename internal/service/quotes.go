package service

import (
	"brandscout-test-task/internal/models"
	"context"
)

type QuotesRepository interface {
	AddQuote(quote *models.Quote)
	GetQuotesByAuthor(author string) ([]*models.Quote, error)
	GetQuote(id uint64) (*models.Quote, error)
	DeleteQuote(id uint64) error
	GetRandomQuote() (*models.Quote, error)
	GetAllQuotes() ([]*models.Quote, error)
}

type QuotesService struct {
	repo QuotesRepository
}

func New(repo QuotesRepository) *QuotesService {
	return &QuotesService{
		repo: repo,
	}
}

func (s *QuotesService) AddQuote(ctx context.Context, quote *models.Quote) {
	s.repo.AddQuote(quote)
}

func (s *QuotesService) GetQuote(ctx context.Context, id uint64) (*models.Quote, error) {
	return s.repo.GetQuote(id)
}

func (s *QuotesService) GetAllQuotes(ctx context.Context, author string) ([]*models.Quote, error) {
	if author == "" {
		return s.repo.GetAllQuotes()
	}

	return s.repo.GetQuotesByAuthor(author)
}

func (s *QuotesService) GetRandomQuote(ctx context.Context) (*models.Quote, error) {
	return s.repo.GetRandomQuote()
}

func (s *QuotesService) DeleteQuote(id uint64) error {
	return s.repo.DeleteQuote(id)
}
