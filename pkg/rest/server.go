package rest

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           router(),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Run() (runerr error) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			switch err {
			case http.ErrServerClosed:
			default:
				runerr = err
			}
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil

}
