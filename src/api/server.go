package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/api/router"
)

type APIServer struct {
	listenAddr string
	store      database.Storage
}

func NewApiServer(listenAddr string, store database.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (srv *APIServer) Run() {
	router := router.New()

	fmt.Printf("\nListening [::]:%s \n", srv.listenAddr)
	log.Fatal(http.ListenAndServe(srv.listenAddr, router))
}
