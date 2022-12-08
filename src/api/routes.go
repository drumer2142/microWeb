package api

import (
	"log"
	"net/http"

	// if middleware put import here

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Controller   func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

func LoadRoutes() []Route {
	routes := APIRoutes //its the array of routes in the same dir
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	// loop the array of routes and shape the HandleFunc
	for _, route := range LoadRoutes() {
		if route.AuthRequired {
			log.Printf("Does not Support Auth middleware yet !!!")
		} else {
			r.HandleFunc(route.URI, route.Controller).Methods(route.Method)
		}
	}

	return r
}
