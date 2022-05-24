package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	_ "github.com/the-bogdan/go-rest-api/app/docs"
	"github.com/the-bogdan/go-rest-api/app/pkg/logging"
)

type App struct {
	logger     logging.Logger
	router     *httprouter.Router
	httpServer *http.Server
	host       string
	port       string
}

func NewApp(host, port string, logger logging.Logger) (App, error) {
	logger.Info("router initializing")

	router := httprouter.New()
	registerHandlers(router)

	return App{
		logger: logger,
		router: router,
		host:   host,
		port:   port,
	}, nil
}

func (a *App) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.host, a.port))
	if err != nil {
		a.logger.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedOrigins:     []string{"http://localhost:3000"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Access-Token"},
		Debug:              false,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	a.logger.Infof("bind application to host http://%s:%s", a.host, a.port)

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")
		default:
			a.logger.Fatal(err)
		}
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Fatal(err)
	}
}
