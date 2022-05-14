package main

import (
	"github.com/the-bogdan/go-rest-api/app/internal/config"
	"github.com/the-bogdan/go-rest-api/app/internal/server"
	"github.com/the-bogdan/go-rest-api/app/pkg/logging"
)

func main() {
	cfg := config.GetConfig()

	logger := logging.GetLogger(cfg.LogLevel)

	app, err := server.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
	app.Run()
}
