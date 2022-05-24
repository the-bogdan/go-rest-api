package main

import (
	"github.com/the-bogdan/go-rest-api/app/internal/server"
	"github.com/the-bogdan/go-rest-api/app/pkg/config"
	"github.com/the-bogdan/go-rest-api/app/pkg/logging"
)

type Config struct {
	IsProd   bool   `env:"IS_PROD" env-default:"false"`
	LogLevel string `env:"LOG_LEVEL" env-default:"info"`

	Listen struct {
		Host string `env:"HOST" env-default:"127.0.0.1"`
		Port string `env:"PORT" env-default:"9990"`
	}
	Postgres struct{}
}

func main() {
	cfg := &Config{}
	config.ReadEnv(cfg)

	logger := logging.Get(cfg.LogLevel)

	logger.WithFields(map[string]interface{}{
		"IS_PROD":   cfg.IsProd,
		"LOG_LEVEL": cfg.LogLevel,
	}).Info("starting HTTP server")

	app, err := server.NewApp(cfg.Listen.Host, cfg.Listen.Port, logger)
	if err != nil {
		logger.Fatal(err)
	}
	app.Run()
}
