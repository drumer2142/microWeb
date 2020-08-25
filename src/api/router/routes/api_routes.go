package routes

import (
  "net/http"
  "github.com/drumer2142/microWeb/src/api/controllers"
)

var api_routes = []Route{
  Route{
    URI: "/all-sites",
    Method: http.MethodGet,
    Controller: controllers.AllSites,
    AuthRequired: false,
  },
  Route{
    URI: "/get/site",
    Method: http.MethodPost,
    Controller: controllers.SiteByDomain,
    AuthRequired: false,
  },
  Route{
    URI: "/store/site",
    Method: http.MethodPost,
    Controller: controllers.StoreSite,
    AuthRequired: false,
  },
}
