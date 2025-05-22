package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		endTime := time.Now()

		log.Printf("Request: %s %s, Duration: %v", r.Method, r.URL.Path, endTime.Sub(startTime))
	}
}
