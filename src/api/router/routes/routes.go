package routes

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

func LoadRoutes() []Route {
	return api_routes
}

func wrapAuth(authRequired bool, h http.HandlerFunc) http.HandlerFunc {
	if !authRequired {
		return h
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// Placeholder: add JWT/session checks here; for now requests still reach the handler.
		log.Printf("auth middleware stub for %s %s", r.Method, r.URL.Path)
		h(w, r)
	}
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range LoadRoutes() {
		h := wrapAuth(route.AuthRequired, route.Controller)
		r.HandleFunc(route.URI, h).Methods(route.Method)
	}
	return r
}
