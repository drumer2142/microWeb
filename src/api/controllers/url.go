package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/api/handler"
	"github.com/drumer2142/microWeb/src/api/models"
	"github.com/drumer2142/microWeb/src/api/repository/repo"
)

func AllSites(w http.ResponseWriter, r *http.Request) {
	db := database.Get()
	if db == nil {
		handler.ResponseError(w, http.StatusInternalServerError, "database not available")
		return
	}
	urlRepository := repo.NewUrlRepo(db)
	sites, err := urlRepository.FindAll()
	if err != nil {
		handler.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.ResponseJSON(w, http.StatusOK, sites)
}

func SiteByDomain(w http.ResponseWriter, r *http.Request) {
	domain := r.FormValue("domain")
	if strings.TrimSpace(domain) == "" {
		handler.ResponseError(w, http.StatusBadRequest, "domain is required")
		return
	}
	db := database.Get()
	if db == nil {
		handler.ResponseError(w, http.StatusInternalServerError, "database not available")
		return
	}
	urlRepository := repo.NewUrlRepo(db)
	sitebydomain, err := urlRepository.FindByDomain(domain)
	if err != nil {
		handler.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(sitebydomain) == 0 {
		handler.ResponseJSON(w, http.StatusNotFound, map[string]string{"message": "No domain found"})
		return
	}
	handler.ResponseJSON(w, http.StatusOK, sitebydomain)
}

func StoreSite(w http.ResponseWriter, r *http.Request) {
	rawURL := r.FormValue("url")
	description := r.FormValue("description")
	if strings.TrimSpace(rawURL) == "" {
		handler.ResponseError(w, http.StatusBadRequest, "url is required")
		return
	}
	domain, err := domainLabelFromURL(rawURL)
	if err != nil {
		handler.ResponseError(w, http.StatusBadRequest, "invalid url")
		return
	}
	site := models.Site{
		URL:         rawURL,
		Domain:      domain,
		Description: description,
	}
	db := database.Get()
	if db == nil {
		handler.ResponseError(w, http.StatusInternalServerError, "database not available")
		return
	}
	if err := db.Create(&site).Error; err != nil {
		handler.ResponseError(w, http.StatusInternalServerError, "could not store website record")
		return
	}
	resp := models.Response{
		Data: models.Data{
			SuccessMsg: "Site Stored Succesfully.",
		},
	}
	handler.ResponseJSON(w, http.StatusOK, &resp)
}

// domainLabelFromURL returns a short label used for matching (first label of the host, without leading "www.").
func domainLabelFromURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if u.Host == "" && !strings.Contains(raw, "://") {
		u, err = url.Parse("https://" + raw)
		if err != nil {
			return "", err
		}
	}
	if u.Host == "" {
		return "", errors.New("missing host")
	}
	host := strings.ToLower(u.Hostname())
	if host == "" {
		return "", errors.New("empty host")
	}
	if strings.HasPrefix(host, "www.") {
		host = strings.TrimPrefix(host, "www.")
	}
	if i := strings.Index(host, "."); i != -1 {
		return host[:i], nil
	}
	return host, nil
}
