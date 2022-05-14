package users

import (
	"net/http"

	"github.com/the-bogdan/go-rest-api/app/internal"
)

const URL = "/users/"

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
	router.HandlerFunc(http.MethodGet, URL+"id/", h.getById)

	// Create instances
	router.HandlerFunc(http.MethodPost, URL, h.create)

	// Update instances
	router.HandlerFunc(http.MethodPut, URL+"id/", h.update)

	// Delete instances
	router.HandlerFunc(http.MethodDelete, URL+"id/", h.delete)
}

// getById
// @Summary Hello simple request returns "Hello world"
// @Tags Test
// @Success 200 "Hello World!!!"
// @Failure 400
// @Router /user [get]
func (h *handler) getById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("{getById}"))
}

func (h *handler) getList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("[getList]"))
}

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("[create]"))
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("[update]"))
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}
