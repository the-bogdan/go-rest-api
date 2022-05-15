package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
// @Summary get user by id
// @Tags Users
// @Success 200
// @Failure 400
// @Router /users/:id [get]
func (h *handler) getById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName(IDENTIFIER))

	if err != nil {
		msg := ErrorMsg{
			Msg: fmt.Sprintf("wrong identifier %s. must be int", params.ByName(IDENTIFIER)),
			Err: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(msg)
		w.Write(body)
	}

	u := User{
		Id:         id,
		FirstName:  "Ermek",
		LastName:   "Pidor",
		MiddleName: "Пиздатый",
		Age:        64,
		IsMale:     true,
		Status:     "Бля заебался курить и в бистро сидеть",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(u)
	w.Write(body)
}

// getList
// @Summary get users list
// @Tags Users
// @Success 200
// @Failure 400
// @Router /users [get]
func (h *handler) getList(w http.ResponseWriter, r *http.Request) {
	usersList := []User{
		{
			Id:         0,
			FirstName:  "Богдан",
			LastName:   "Паршинцев",
			MiddleName: "Backend",
			Age:        70,
			IsMale:     true,
			Status:     "Че покурим?",
		},
		{
			Id:         1,
			FirstName:  "Ermek",
			LastName:   "Pidor",
			MiddleName: "Пиздатый",
			Age:        64,
			IsMale:     true,
			Status:     "Бля заебался курить и в бистро сидеть",
		},
		{
			Id:        2,
			FirstName: "Aidar",
			LastName:  "Slojnoe",
			Age:       16,
			IsMale:    true,
			Status:    "Вырубился через 10сек после того, как лёг",
		},
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(usersList)
	w.Write(body)
}

// create
// @Summary create user
// @Tags Users
// @Success 201
// @Failure 400
// @Router /users [post]
func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("[create]"))
}

// update
// @Summary update user info by user id
// @Tags Users
// @Success 200
// @Failure 400
// @Router /users/:id [put]
func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("[update %s]", params.ByName(IDENTIFIER))))
}

// delete
// @Summary delete user by id
// @Tags Users
// @Success 204
// @Failure 400
// @Router /users/:id [delete]
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	//params := httprouter.ParamsFromContext(r.Context())

	w.WriteHeader(http.StatusNoContent)
}
