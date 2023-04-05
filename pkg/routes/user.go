package routes

import (
	"otomo_golang/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoute = func(router *mux.Router) {
	router.HandleFunc("/api/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/user/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/delete/{user_id}", controllers.DeleteUser).Methods(("DELETE"))
	router.HandleFunc("/api/user", controllers.ListUsers).Methods("GET")
}
