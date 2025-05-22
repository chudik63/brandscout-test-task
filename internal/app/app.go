package app

import (
	"brandscout-test-task/internal/config"
	"brandscout-test-task/internal/limiter"
	"brandscout-test-task/internal/repository"
	"brandscout-test-task/internal/server"
	"brandscout-test-task/internal/service"
	"brandscout-test-task/internal/transport/http/handler"
	"brandscout-test-task/internal/transport/http/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTime = 5 * time.Second

func Run() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	mux := http.NewServeMux()

	repo := repository.New()

	service := service.New(repo)

	handler := handler.New(service)

	limiter := limiter.New(cfg.RateLimit)

	routes.RegistrateRoutes(handler, limiter, mux)

	server := server.NewServer(cfg.HTTP, mux)

	go func() {
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited gracefully")
}
