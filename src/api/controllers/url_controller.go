package controllers


import (
  _"log"
  "strings"
  "net/http"
  "github.com/drumer2142/microWeb/src/api/handler"
  "github.com/drumer2142/microWeb/src/api/models"
  "github.com/drumer2142/microWeb/src/api/repository"
  "github.com/drumer2142/microWeb/src/api/repository/repo"
  "github.com/drumer2142/microWeb/src/api/database"
)

func AllSites(w http.ResponseWriter, r *http.Request){
  db, err := database.Connect()
  if err != nil {
		handler.ResponseJSON(w, http.StatusInternalServerError, err)
		return
  }
  repo := repo.NewUrlRepo(db)
  func(urlRepository repository.UrlRepository){
    sites, err := urlRepository.FindAll()
    if err != nil {
      handler.ResponseError(w, http.StatusNotFound, "No result found")
      return
    }
    handler.ResponseJSON(w, http.StatusOK, sites)
  }(repo)
}

func SiteByDomain(w http.ResponseWriter, r *http.Request){
    domain := r.FormValue("domain")
    db, _ := database.Connect()
    sitebydomain := &models.Site{}
    err := db.Find(sitebydomain, "domain = ?", domain).RecordNotFound()
    if err == true {
        handler.ResponseError(w, http.StatusNotFound, "No result found")
        return
    }
    handler.ResponseJSON(w, http.StatusOK, sitebydomain)
}

func StoreSite(w http.ResponseWriter, r *http.Request){
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
    var site = models.Site{
      URL: url, 
      Domain: domain,
      Description: description,
    }

    db, _ := database.Connect()
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
