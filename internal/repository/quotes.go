package repository

import (
	"brandscout-test-task/internal/models"
	"sync"

	"math/rand/v2"
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

func (r *QuotesRepository) GetRandomQuote() *models.Quote {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.quotes) == 0 {
		return nil
	}

	randID := rand.IntN(len(r.quotes))

	return r.quotes[uint64(randID)]
}

func (r *QuotesRepository) GetQuotesByAuthor(author string) []*models.Quote {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quotes := make([]*models.Quote, 0)
	for _, quote := range r.quotes {
		if quote.Author == author {
			quotes = append(quotes, quote)
		}
	}

	return quotes
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
