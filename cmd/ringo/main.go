package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	os.Exit(run())
}

func run() (code int) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"alive": true,
		})
	})

	server := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			code = 1
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		switch err {
		case http.ErrServerClosed:
		default:
			return 1
		}
	}

	return 0
}
