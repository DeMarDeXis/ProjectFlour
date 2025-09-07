package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

const (
	DBName = "postgres"
)

// TODO: was changed
type StorageConfig struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Username string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	DBName   string `yaml:"dbname" env:"DB_NAME" env-default:"postgres"`
	SSLMode  string `yaml:"ssl_mode" env:"DB_SSLMODE" env-default:"disable"`
}

func New(cfg StorageConfig, logg *slog.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open(DBName, builderConnection(cfg))
	if err != nil {
		logg.Error("failed to ping db", slog.Any("error", err), slog.String("connection_string", builderConnection(cfg)))
		return nil, fmt.Errorf("db.Ping - %w", err)
	}

	err = db.Ping()
	if err != nil {
		logg.Error("failed to ping db", slog.Any("error", err), slog.String("connection_string", builderConnection(cfg)))
		return nil, fmt.Errorf("db.Ping - %w", err)
	}

	return db, nil
}

func Stop(db *sqlx.DB, logg *slog.Logger) error {
	const op = "storage.postgres.Stop"

	if db == nil {
		logg.Error("db is nil", slog.String("op", op))
		return fmt.Errorf("db is nil")
	}

	logg.Debug("close db connection", slog.String("op", op))

	if err := db.Close(); err != nil {
		logg.Error("failed to close db connection", slog.String("op", op), slog.Any("error", err))
		return fmt.Errorf("db.Close - %w", err)
	}

	logg.Debug("db connection closed successfully", slog.String("op", op))
	return nil
}

func builderConnection(cfg StorageConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
}
