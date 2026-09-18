package users

import (
	"api-social-media/internal/core/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", middlewares.Authenticate(CreateUser)).Methods(http.MethodPost)
	router.HandleFunc("/users", middlewares.Authenticate(GetUsers)).Methods(http.MethodGet)
	router.HandleFunc("/users/", middlewares.Authenticate(GetUsersByNickOrName)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", middlewares.Authenticate(GetUser)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", middlewares.Authenticate(UpdateUser)).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", middlewares.Authenticate(DeleteUser)).Methods(http.MethodDelete)
}
