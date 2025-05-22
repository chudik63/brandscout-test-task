package limiter

import (
	"brandscout-test-task/internal/config"
	"net/http"

	"golang.org/x/time/rate"
)

type Limiter struct {
	*rate.Limiter
}

func New(cfg config.RateLimiterConfig) *Limiter {
	return &Limiter{rate.NewLimiter(rate.Limit(cfg.Limit), cfg.Limit)}
}

func (l *Limiter) RateLimitMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !l.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}
