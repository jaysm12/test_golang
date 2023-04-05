package main

import (
	"fmt"
	"log"
	"net/http"
	"otomo_golang/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUserRoute(r)
	http.Handle("/", r)

	fmt.Println("Server running on port: 3000")
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
