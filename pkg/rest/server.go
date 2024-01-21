package rest

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rinchsan/ringo/pkg/zlog"
)

type Server struct {
	httpServer *http.Server
	logger     *zlog.Logger
}

func NewServer(logger *zlog.Logger) *Server {
	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           router(logger),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	return &Server{
		httpServer: httpServer,
		logger:     logger,
	}
}

func (s *Server) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error)
	go func() {
		s.logger.Info("server started")
		if err := s.httpServer.ListenAndServe(); err != nil {
			switch err {
			case http.ErrServerClosed:
			default:
				errCh <- err
			}
		}
	}()

	select {
	case <-ctx.Done():
		s.logger.Info("graceful shutdown started")
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := s.httpServer.Shutdown(ctx); err != nil {
			return err
		}
	case err := <-errCh:
		s.logger.Error(err)
		return err
	}

	return nil
}
