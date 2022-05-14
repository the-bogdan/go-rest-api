package server

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/the-bogdan/go-rest-api/app/internal"
	"github.com/the-bogdan/go-rest-api/app/internal/handlers/hello-world"
	"github.com/the-bogdan/go-rest-api/app/internal/handlers/users"
)

func registerHandlers(router internal.Router) {
	// swagger docs initializing
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	hello_world.NewHandler().Register(router)
	users.NewHandler().Register(router)
}
