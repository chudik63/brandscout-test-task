package handler

import (
	"brandscout-test-task/internal/models"
	"brandscout-test-task/internal/repository"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const requestTimeout = 10 * time.Second

//go:generate go run github.com/vektra/mockery/v2@latest --name Service
type Service interface {
	AddQuote(ctx context.Context, quote *models.Quote)
	GetQuote(ctx context.Context, id uint64) (*models.Quote, error)
	GetAllQuotes(ctx context.Context, author string) ([]*models.Quote, error)
	GetRandomQuote(ctx context.Context) (*models.Quote, error)
	DeleteQuote(ctx context.Context, id uint64) error
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddQuote(w http.ResponseWriter, r *http.Request) {
	var quote models.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	h.service.AddQuote(ctx, &quote)

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	quotes, err := h.service.GetAllQuotes(ctx, author)
	if err != nil {
		log.Println("Failed to get quotes:", err)
		http.Error(w, "Failed to retrieve quotes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	quote, err := h.service.GetRandomQuote(ctx)
	if err != nil {
		log.Println("Failed to get random quote:", err)
		http.Error(w, "Failed to retrieve quote", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func (h *Handler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/quotes/")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	if err := h.service.DeleteQuote(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "Quote not found", http.StatusNotFound)
			return
		}

		log.Println("Failed to delete quote:", err)
		http.Error(w, "Failed to delete quote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
