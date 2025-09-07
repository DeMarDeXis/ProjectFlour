package httpServer

import (
	"ProjectFlour/pkg/config"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppServer struct {
	log    *slog.Logger
	server *http.Server
}

func New(log *slog.Logger, cfg config.HTTPServer, handler http.Handler) *AppServer {
	return &AppServer{
		log: log,
		server: &http.Server{
			Addr:         cfg.Host + ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

func (s *AppServer) Start() error {
	const op = "app.httpServer.Start"

	s.log.Info("Trying to start http server", slog.String("op", op))

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error("Failed to start http server", slog.String("op", op), slog.String("error", err.Error()))
		}
	}()

	return nil
}

func (s *AppServer) Stop() error {
	const op = "app.httpServer.Stop"

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	s.log.Info("Shutting down server...", slog.String("op", op))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Error("Failed to stop server", slog.String("op", op), slog.String("error", err.Error()))
		return err
	}

	s.log.Info("Server stopped", slog.String("op", op))
	return nil
}
