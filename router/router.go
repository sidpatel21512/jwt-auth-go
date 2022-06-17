package router

import (
	"github.com/gorilla/mux"
	"github.com/sidpatel21512/jwt-auth-go/controller"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/register", controller.RegisterUserHandler).Methods("POST")
	subRouter.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	subRouter.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")
	subRouter.HandleFunc("/users", controller.UsersHandler).Methods("GET")
	subRouter.HandleFunc("/healthcheck", controller.HealthCheckHandler).Methods("GET")

	return router
}
