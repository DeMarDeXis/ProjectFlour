package wsServer

import (
	"ProjectFlour/internal/handlers/webSocketHandler"
	"ProjectFlour/pkg/config"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type WsServer struct {
	log     *slog.Logger
	httpSrv *http.Server
	handler *webSocketHandler.WebSocketHandler
}

func New(log *slog.Logger, cfg config.WebSocket, handler *webSocketHandler.WebSocketHandler) *WsServer {
	return &WsServer{
		log: log,
		httpSrv: &http.Server{
			Addr:    cfg.Host + ":" + cfg.Port,
			Handler: http.HandlerFunc(handler.HandleConnections),
		},
		handler: handler,
	}
}

func (s *WsServer) Start() error {
	const op = "app.wsServer.Start"

	s.log.Info("Trying to start ws server", slog.String("op", op), slog.String("address", s.httpSrv.Addr))

	go func() {
		if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error("Failed to start ws server", slog.String("op", op), slog.String("error", err.Error()))
		}
	}()

	return nil
}

func (s *WsServer) Stop() error {
	const op = "app.wsServer.Stop"

	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	//
	//<-quit

	s.log.Info("Shutting down server...", slog.String("op", op))

	if s.handler != nil {
		s.handler.DisconnectAll()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.log.Error("Failed to stop WebSocket server", slog.String("op", op), slog.String("error", err.Error()))
		return err
	}

	s.log.Info("WebSocket server stopped", slog.String("op", op))
	return nil
}
