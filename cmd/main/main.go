package main

import (
	"log"
	"net/http"
	"otomo_golang/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUserRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
