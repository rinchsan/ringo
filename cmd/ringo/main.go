package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "go.uber.org/automaxprocs"
)

type ringo struct {
	REST CmdREST `cmd:"" help:"Run REST Server"`
}

type CmdREST struct {
}

func (c *CmdREST) Run() (runerr error) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := http.Server{
		Addr:              ":8080",
		Handler:           router(),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			runerr = err
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		switch err {
		case http.ErrServerClosed:
		default:
			return err
		}
	}

	return nil
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
		return http.MaxBytesHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		}), 4096)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"alive": true,
		})
	})

	return r
}

func main() {
	var ringo ringo
	ctx := kong.Parse(&ringo)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
