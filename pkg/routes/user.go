package routes

import (
	"otomo_golang/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoute = func(router *mux.Router) {
	// router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.ListUsers).Methods("GET")
	// router.HandleFunc("/login", controllers.Login).Methods("POST")
	// router.HandleFunc("/user", controllers.Deleteuser).Methods(("DELETE"))
}
