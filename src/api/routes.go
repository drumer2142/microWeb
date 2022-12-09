package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Controller   func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

func (srv *APIServer) SetupRoutes(r *mux.Router) *mux.Router {
	// loop the array of routes and shape the HandleFunc
	for _, route := range srv.LoadRoutes() {
		if route.AuthRequired {
			log.Printf("Does not Support Auth middleware yet !!!")
		} else {
			r.HandleFunc(route.URI, route.Controller).Methods(route.Method)
		}
	}

	return r
}
