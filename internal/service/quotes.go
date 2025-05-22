package service

import (
	"brandscout-test-task/internal/models"
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name QuotesRepository
type QuotesRepository interface {
	AddQuote(ctx context.Context, quote *models.Quote)
	GetQuotesByAuthor(ctx context.Context, author string) ([]*models.Quote, error)
	GetQuote(ctx context.Context, id uint64) (*models.Quote, error)
	DeleteQuote(ctx context.Context, id uint64) error
	GetRandomQuote(ctx context.Context) (*models.Quote, error)
	GetAllQuotes(ctx context.Context) ([]*models.Quote, error)
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
	s.repo.AddQuote(ctx, quote)
}

func (s *QuotesService) GetQuote(ctx context.Context, id uint64) (*models.Quote, error) {
	return s.repo.GetQuote(ctx, id)
}

func (s *QuotesService) GetAllQuotes(ctx context.Context, author string) ([]*models.Quote, error) {
	if author == "" {
		return s.repo.GetAllQuotes(ctx)
	}

	return s.repo.GetQuotesByAuthor(ctx, author)
}

func (s *QuotesService) GetRandomQuote(ctx context.Context) (*models.Quote, error) {
	return s.repo.GetRandomQuote(ctx)
}

func (s *QuotesService) DeleteQuote(ctx context.Context, id uint64) error {
	return s.repo.DeleteQuote(ctx, id)
}
