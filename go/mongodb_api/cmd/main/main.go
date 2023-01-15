package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vyrekxd/mongodb_api/pkg/routes"
)

func main() {

	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
