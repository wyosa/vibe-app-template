package app

import (
	"context"
	"net/http"
	"time"

	"app/api/internal/delivery/http/gen"
	"app/api/internal/delivery/http/handlers"
	"app/api/internal/delivery/http/middlewares"
	"app/api/internal/delivery/http/server"
)

type App struct {
	container *Container
	server    *server.Server
}

func New(container *Container) *App {
	healthHandler := handlers.NewHealthHandler()
	strictHandler := gen.NewStrictHandler(healthHandler, nil)
	mux := gen.HandlerFromMux(strictHandler, http.NewServeMux())

	handler := middlewares.Recovery(container.Log)(
		middlewares.RequestID(
			middlewares.Logging(container.Log)(mux),
		),
	)

	srv := server.New(handler, container.Config.HTTPPort, container.Log)

	return &App{
		container: container,
		server:    srv,
	}
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := a.server.Start(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return a.server.Shutdown(shutdownCtx)
	case err := <-errCh:
		return err
	}
}
