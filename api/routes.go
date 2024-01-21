package api

import (
	"src/service/handler"
	"src/service/test"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Define a route that will handle concurrent requests
	r.HandleFunc("/concurrent", test.ConcurrentHandler).Methods("GET")
	r.HandleFunc("/search", handler.SearchHandler).Methods("GET")
	return r
}
