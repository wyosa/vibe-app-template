package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	log        *slog.Logger
}

func New(handler http.Handler, port int64, log *slog.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		log: log,
	}
}

func (s *Server) Start() error {
	s.log.Info("server starting", "addr", s.httpServer.Addr)

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("server shutting down")

	return s.httpServer.Shutdown(ctx)
}
