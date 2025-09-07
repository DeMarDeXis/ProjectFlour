package main

import (
	_ "ProjectFlour/docs"
	"ProjectFlour/internal/app/httpServer"
	"ProjectFlour/internal/app/wsServer"
	"ProjectFlour/internal/events"
	"ProjectFlour/internal/handlers/httpHandler"
	"ProjectFlour/internal/handlers/webSocketHandler"
	"ProjectFlour/internal/service"
	"ProjectFlour/internal/storage"
	"ProjectFlour/internal/storage/postgres"
	"ProjectFlour/pkg/config"
	"ProjectFlour/pkg/lib/logger/handler/slogpretty"
	"flag"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// parse flags
	var configPath = flag.String("config", "", "Path to config file (e.g. -config=config/config.yaml)")
	flag.Parse()

	// init config
	cfg := config.InitConfig(*configPath)

	// init logger
	logg := setupLogger(cfg.Env)
	logg.Info("Config path", slog.String("config", *configPath))
	logg.Info("starting project flour", slog.String("env", cfg.Env))

	// init db
	db, err := postgres.New(postgres.StorageConfig{
		Host:     cfg.StorageConfig.Host,
		Port:     cfg.StorageConfig.Port,
		Username: cfg.StorageConfig.Username,
		Password: cfg.StorageConfig.Password,
		DBName:   cfg.StorageConfig.DBName,
		SSLMode:  cfg.StorageConfig.SSLMode,
	}, logg)
	if err != nil {
		logg.Error("failed to init db", slog.Any("error", err.Error()))
		os.Exit(1)
	}

	logg.Debug("init db",
		slog.String("host", cfg.StorageConfig.Host),
		slog.String("port", cfg.StorageConfig.Port),
		slog.String("username", cfg.StorageConfig.Username),
		slog.String("dbname", cfg.StorageConfig.DBName),
		slog.String("sslmode", cfg.StorageConfig.SSLMode),
		slog.String("env", cfg.Env))

	strg := storage.NewStorage(db)

	//init eventBus
	eventBus := events.NewEventBus()

	// init service
	srvc := service.New(strg, eventBus)

	// init http httpHandler
	httpHndlr := httpHandler.NewHTTTPHandler(srvc, logg)

	logg.Debug("cfg",
		slog.Any("allowed_origins", cfg.CORS.AllowedOrigins),
	)

	// add routes
	routesToSrv := httpHndlr.InitRoutes(logg)

	handlerWithCors := cors.New(cors.Options{
		AllowedOrigins:   cfg.CORS.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            cfg.Env == envLocal,
	}).Handler(routesToSrv)

	// init server
	httpSRV := httpServer.New(logg, cfg.HTTPServer, handlerWithCors) // Instead of routesToSкм
	//httpSRV := httpServer.New(logg, cfg.HTTPServer, routesToSrv) // TODO: change it

	// add websocket handlers
	wsHanlers := webSocketHandler.New(logg, eventBus)

	// init websocket server
	wsSRV := wsServer.New(logg, cfg.WebSocket, wsHanlers)

	go func() {
		if err := wsSRV.Start(); err != nil {
			logg.Error("failed to start ws server", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	if err := httpSRV.Start(); err != nil {
		logg.Error("failed to start http server", slog.Any("error", err))
		os.Exit(1)
	}

	// Обработка graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logg.Info("Shutting down servers...")

	// init graceful shutdown
	if err := httpSRV.Stop(); err != nil {
		logg.Error("failed to stop http server", slog.Any("error", err))
		os.Exit(1)
	}

	if err := wsSRV.Stop(); err != nil {
		logg.Error("failed to stop ws server", slog.Any("error", err))
		os.Exit(1)
	}

	// stop db
	if err := postgres.Stop(db, logg); err != nil {
		logg.Error("failed to stop db", slog.Any("error", err))
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = initSlogPretty()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func initSlogPretty() *slog.Logger {
	opts := slogpretty.PrettyHandlersOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handl := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handl)
}

// TODO: i need to fill all env values in current remote GitHub
// TODO: make template's downloader for excel files [12.08.2025] [MAIN] - Now it's under development
// TODO: mockery business must be correct <-- It's in progress...
//		packages:
//			ProjectFlour/internal/storage:
//				config:
//					include-interface-regex: '^AuthorizationStorage$'

// TODO[DB_Prop]:
// 	ssl_mode: "" <- fix it

// TODO[FUTURE]:
// 	custom_logging_outputer:
// 		level: "info"
//  	format: "json"
// 		file_path: "/var/log/app.log"

// TODO[FUTURE]:
// 	monitoring:
// 		prometheus_enabled: true
// 		metrics_port: 9090

// TODO[FUTURE]:
// 	rate_limiting:
// 		enabled: true
// 		requests_per_minute: 1000
