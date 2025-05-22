package repository

import (
	"brandscout-test-task/internal/models"
	"context"
	"errors"
	"sync"

	"math/rand/v2"
)

var (
	ErrNotFound = errors.New("nothing was found")
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

func (r *QuotesRepository) AddQuote(ctx context.Context, quote *models.Quote) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uint64(len(r.quotes) + 1)
	quote.ID = id

	r.quotes[id] = quote
}

func (r *QuotesRepository) GetRandomQuote(ctx context.Context) (*models.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.quotes) == 0 {
		return nil, ErrNotFound
	}

	randID := rand.IntN(len(r.quotes))

	return r.quotes[uint64(randID)], nil
}

func (r *QuotesRepository) GetQuotesByAuthor(ctx context.Context, author string) ([]*models.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quotes := make([]*models.Quote, 0)
	for _, quote := range r.quotes {
		if quote.Author == author {
			quotes = append(quotes, quote)
		}
	}

	if len(quotes) == 0 {
		return nil, ErrNotFound
	}

	return quotes, nil
}

func (r *QuotesRepository) GetQuote(ctx context.Context, id uint64) (*models.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quote, exists := r.quotes[id]
	if !exists {
		return nil, ErrNotFound
	}

	return quote, nil
}

func (r *QuotesRepository) DeleteQuote(ctx context.Context, id uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.quotes[id]

	if !exists {
		return ErrNotFound
	}

	delete(r.quotes, id)

	return nil
}

func (r *QuotesRepository) GetAllQuotes(ctx context.Context) ([]*models.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.quotes) == 0 {
		return nil, ErrNotFound
	}

	quotes := make([]*models.Quote, 0, len(r.quotes))
	for _, quote := range r.quotes {
		quotes = append(quotes, quote)
	}

	return quotes, nil
}
