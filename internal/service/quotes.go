package service

import "brandscout-test-task/internal/models"

type QuotesRepository interface {
	AddQuote(quote *models.Quote)
	GetQuote(id uint64) (*models.Quote, bool)
	DeleteQuote(id uint64)
	GetAllQuotes() []*models.Quote
}

type QuotesService struct {
	repo QuotesRepository
}

func New(repo QuotesRepository) *QuotesService {
	return &QuotesService{
		repo: repo,
	}
}

func (s *QuotesService) AddQuote(quote *models.Quote) error {

}

func (s *QuotesService) GetQuote(id uint64) error {

}

func (s *QuotesService) GetAllQuotes() ([]*models.Quote, error) {

}

func (s *QuotesService) GetRandomQuote() (*models.Quote, error) {

}

func (s *QuotesService) DeleteQuote(id uint64) error {

}
