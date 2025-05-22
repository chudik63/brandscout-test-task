package routes

import (
	"brandscout-test-task/internal/transport/http/middleware"
	"net/http"
)

type Handler interface {
	CreateQuote(w http.ResponseWriter, r *http.Request)
	GetQuotes(w http.ResponseWriter, r *http.Request)
	GetRandomQuote(w http.ResponseWriter, r *http.Request)
	DeleteQuote(w http.ResponseWriter, r *http.Request)
}

type Limiter interface {
	RateLimitMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)
}

func RegistrateRoutes(h Handler, limiter Limiter, mux *http.ServeMux) {
	mux.HandleFunc("POST /quotes", limiter.RateLimitMiddleware(middleware.LoggingMiddleware(h.CreateQuote)))

	mux.HandleFunc("GET /quotes", limiter.RateLimitMiddleware(middleware.LoggingMiddleware(h.GetQuotes)))

	mux.HandleFunc("GET /quotes/random", limiter.RateLimitMiddleware(middleware.LoggingMiddleware(h.GetRandomQuote)))

	mux.HandleFunc("DELETE /quotes/:id", limiter.RateLimitMiddleware(middleware.LoggingMiddleware(h.DeleteQuote)))
}
