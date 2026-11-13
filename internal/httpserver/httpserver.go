package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"streamgate/internal/api"
	"streamgate/internal/dependencies"
	"time"
)

type Server struct {
	srv    *http.Server
	logger *slog.Logger
}

func New(deps dependencies.Service) *Server {
	return &Server{
		srv: &http.Server{
			Addr:              deps.Config().ListenAddr,
			Handler:           api.New(deps),
			ReadHeaderTimeout: time.Second * 10,
			IdleTimeout:       time.Minute * 2,
		},
		logger: deps.Logger(),
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("Starting HTTP server", "addr", s.srv.Addr)
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.srv.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		return err

	case <-ctx.Done():
		s.logger.Info("Shutting down HTTP server")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := s.srv.Shutdown(shutdownCtx); err != nil {
			_ = s.srv.Close()
			return fmt.Errorf("failed to shutdown HTTP server: %w", err)
		}

		return nil
	}
}
