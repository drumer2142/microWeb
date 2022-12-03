package routes

import (
	"net/http"

	"github.com/drumer2142/microWeb/src/api/controllers"
)

var api_routes = []Route{
	{
		URI:          "/all-sites",
		Method:       http.MethodGet,
		Controller:   controllers.AllSites,
		AuthRequired: false,
	},
	{
		URI:          "/get/site",
		Method:       http.MethodPost,
		Controller:   controllers.SiteByDomain,
		AuthRequired: false,
	},
	{
		URI:          "/store/site",
		Method:       http.MethodPost,
		Controller:   controllers.StoreSite,
		AuthRequired: false,
	},
}
