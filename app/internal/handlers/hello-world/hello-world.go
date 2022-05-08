package hello_world

import (
	"net/http"

	"github.com/the-bogdan/go-rest-api/app/internal"
)

const URL = "/hello"

type Handler struct {
}

func (h *Handler) Register(router internal.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Hello)
}

// Hello
// @Summary Hello simple request returns "Hello world"
// @Tags Test
// @Success 200 "Hello World!!!"
// @Failure 400
// @Router /hello [get]
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world!!!"))
}
