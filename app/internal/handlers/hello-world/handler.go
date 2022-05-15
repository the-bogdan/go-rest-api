package hello_world

import (
	"net/http"

	"github.com/the-bogdan/go-rest-api/app/internal"
)

const URL = "/hello"

type handler struct{}

func NewHandler() internal.Handler {
	return &handler{}
}

func (h *handler) Register(router internal.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.hello)
}

// Hello
// @Summary Hello simple request returns "Hello world"
// @Tags Test
// @Success 200 "Hello World!!!"
// @Failure 400
// @Router /hello [get]
func (h *handler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!!!"))
}
