package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/api/handler"
	"github.com/drumer2142/microWeb/src/api/models"
)

type APIServer struct {
	listenAddr string
	store      database.Storage
	routes     Route
}

func NewApiServer(listenAddr string, store database.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (srv *APIServer) Run() {

	router := NewRouter()

	fmt.Printf("\nListening [::]:%s \n", srv.listenAddr)
	log.Fatal(http.ListenAndServe(srv.listenAddr, router))
}

func (srv *APIServer) StoreWebsite(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	description := r.FormValue("description")
	// split the full url with . before the domain name
	url_split1 := strings.Split(url, ".")
	url_split2 := strings.Split(url, "/")
	domain := url_split1[1]
	// get the domain name if www dows not exist in the url
	if url_split2[2][0:3] != "www" {
		url_split3 := strings.Split(url_split2[2], ".")
		domain = url_split3[0]
	}
	var webSite = &models.Website{
		URL:         url,
		Domain:      domain,
		Description: description,
	}

	err := srv.store.CreateWebSite(webSite)

	if err != nil {
		handler.ResponseError(w, http.StatusNotFound, "Can not create website record")
		return
	}

	resp := models.Response{
		Data: models.Data{
			SuccessMsg: "Site Stored Succesfully.",
		},
	}
	handler.ResponseJSON(w, http.StatusOK, &resp)
}

func (srv *APIServer) RetriveAllWebsites(w http.ResponseWriter, r *http.Request) {
	allWebsites, err := srv.store.GetAllWebSites()
	if err != nil {
		handler.ResponseError(w, http.StatusNotFound, "No result found")
		return
	}

	handler.ResponseJSON(w, http.StatusOK, allWebsites)
}

func (srv *APIServer) RetriveWebsiteByDomain(w http.ResponseWriter, r *http.Request) {
	domain := r.FormValue("domain")

	webSiteByDomain, err := srv.store.GetByDomainName(domain)
	if err != nil {
		handler.ResponseJSON(w, http.StatusNotFound, err)
		return
	}

	handler.ResponseJSON(w, http.StatusOK, webSiteByDomain)
}
