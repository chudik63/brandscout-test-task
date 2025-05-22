package server

import (
	"brandscout-test-task/internal/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.HTTPConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	log.Print("Server is running on port", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
