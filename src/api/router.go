package api

import (
	"github.com/gorilla/mux"
)

func (srv *APIServer) NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return srv.SetupRoutes(r)
}
