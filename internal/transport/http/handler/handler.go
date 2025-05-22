package handler

import (
	"net/http"
	"time"
)

const requestTimeout = 10 * time.Second

//go:generate go run github.com/vektra/mockery/v2@latest --name Service
type Service interface {
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a quote
}

func (h *Handler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting quotes
}

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a random quote
}

func (h *Handler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a quote
}
