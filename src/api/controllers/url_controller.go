package controllers

import (
	_ "log"
	"net/http"
	"strings"

	"github.com/drumer2142/microWeb/src/api/handler"
	"github.com/drumer2142/microWeb/src/api/models"
	"github.com/drumer2142/microWeb/src/api/repository"
	"github.com/drumer2142/microWeb/src/api/repository/repo"
)

func AllSites(w http.ResponseWriter, r *http.Request) {

	repo := repo.NewUrlRepo(db)
	func(urlRepository repository.UrlRepository) {
		sites, err := urlRepository.FindAll()
		if err != nil {
			handler.ResponseError(w, http.StatusNotFound, "No result found")
			return
		}
		handler.ResponseJSON(w, http.StatusOK, sites)
	}(repo)
}

func SiteByDomain(w http.ResponseWriter, r *http.Request) {
	domain := r.FormValue("domain")

	repo := repo.NewUrlRepo(db)
	func(urlRepository repository.UrlRepository) {
		sitebydomain, err := urlRepository.FindByDomain(domain)
		if err != nil {
			handler.ResponseJSON(w, http.StatusNotFound, err)
			return
		}
		if len(sitebydomain) == 0 {
			handler.ResponseJSON(w, http.StatusNotFound, "No domain found")
			return
		}
		handler.ResponseJSON(w, http.StatusOK, sitebydomain)
	}(repo)
}

func StoreSite(w http.ResponseWriter, r *http.Request) {
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
	var site = models.Website{
		URL:         url,
		Domain:      domain,
		Description: description,
	}

	err := db.Create(&site).RecordNotFound()
	if err == true {
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
