package config

import (
	"ProjectFlour/internal/storage/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	Env                    string `yaml:"env" env:"APP_ENV" env-default:"local"`
	HTTPServer             `yaml:"http_server"`
	WebSocket              `yaml:"ws_server"`
	postgres.StorageConfig `yaml:"db_postgres"`
	CORS                   CORSConfig `yaml:"cors"`
}
type HTTPServer struct {
	Host         string        `yaml:"address" env-default:"localhost"`
	Port         string        `yaml:"port" env-default:"8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env-default:"30s"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env-default:"10s"`
}

type WebSocket struct {
	Host            string        `yaml:"address" env-default:"localhost"`
	Port            string        `yaml:"port" env-default:"8081"`
	ReadBufferSize  int           `yaml:"read_buffer_size" env-default:"1024"`
	WriteBufferSize int           `yaml:"write_buffer_size" env-default:"1024"`
	PingInterval    time.Duration `yaml:"ping_interval" env-default:"60s"`
}

type CORSConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
}

// InitConfig load config for app.
// main ".env" have priority over ".env.local" forever
func InitConfig(configPath string) *Config {
	var cfg Config

	_ = godotenv.Load(".env")

	envFile := ""
	switch configPath {
	case "./config/local.yaml":
		envFile = ".local.env"
	case "./config/dev.yaml":
		envFile = ".dev.env"

	}

	if envFile != "" {
		if err := godotenv.Overload(envFile); err != nil {
			log.Printf("Warning: failed to load %s: %v", envFile, err)
		}
	}

	if configPath != "" {
		if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
			log.Printf("Achtung: failed to read cfg file: %v", err)
		}
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Error reading environment variables: %v", err)
	}

	return &cfg
}
