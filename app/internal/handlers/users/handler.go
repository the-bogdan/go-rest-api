package users

import (
	"fmt"
	"net/http"

	// TODO remove dependency on httprouter
	"github.com/julienschmidt/httprouter"

	"github.com/the-bogdan/go-rest-api/app/internal"
)

const (
	URL               = "/users"
	IDENTIFIER string = "id"
)

// handler struct for working with users
type handler struct{}

// NewHandler returns current module handler
func NewHandler() internal.Handler {
	return &handler{}
}

// Register add current module handlers to router and register paths
func (h *handler) Register(router internal.Router) {
	// Get instances
	router.HandlerFunc(http.MethodGet, URL, h.getList)
	router.HandlerFunc(http.MethodGet, fmt.Sprintf("%s/:%s", URL, IDENTIFIER), h.getById)

	// Create instances
	router.HandlerFunc(http.MethodPost, URL, h.create)

	// Update instances
	router.HandlerFunc(http.MethodPut, fmt.Sprintf("%s/:%s", URL, IDENTIFIER), h.update)

	// Delete instances
	router.HandlerFunc(http.MethodDelete, fmt.Sprintf("%s/:%s", URL, IDENTIFIER), h.delete)
}

// getById
// @Summary Hello simple request returns "Hello world"
// @Tags Test
// @Success 200 "Hello World!!!"
// @Failure 400
// @Router /user [get]
func (h *handler) getById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{getById %s}", params.ByName(IDENTIFIER))))
}

func (h *handler) getList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[getList]"))
}

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("[create]"))
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("[update %s]", params.ByName(IDENTIFIER))))
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	//params := httprouter.ParamsFromContext(r.Context())

	w.WriteHeader(http.StatusNoContent)
}
