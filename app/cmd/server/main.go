package main

import (
	"github.com/the-bogdan/go-rest-api/app/internal"
	"github.com/the-bogdan/go-rest-api/app/internal/server"
	"github.com/the-bogdan/go-rest-api/app/pkg/config"
	"github.com/the-bogdan/go-rest-api/app/pkg/logging"
)

func main() {

	cfg := &internal.Config{}
	config.GetConfig(cfg)

	logger := logging.GetLogger(cfg.LogLevel)

	app, err := server.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
	app.Run()
}
