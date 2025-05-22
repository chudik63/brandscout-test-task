package repository

import (
	"brandscout-test-task/internal/models"
	"sync"
)

type QuotesRepository struct {
	mu     sync.RWMutex
	quotes map[uint64]*models.Quote
}

func New() *QuotesRepository {
	return &QuotesRepository{
		quotes: make(map[uint64]*models.Quote),
	}
}

func (r *QuotesRepository) AddQuote(quote *models.Quote) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.quotes[quote.ID] = quote
}

func (r *QuotesRepository) GetQuote(id uint64) (*models.Quote, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quote, exists := r.quotes[id]
	return quote, exists
}

func (r *QuotesRepository) DeleteQuote(id uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.quotes, id)
}

func (r *QuotesRepository) GetAllQuotes() []*models.Quote {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quotes := make([]*models.Quote, 0, len(r.quotes))
	for _, quote := range r.quotes {
		quotes = append(quotes, quote)
	}

	return quotes
}
