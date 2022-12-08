package api

import (
	"net/http"
)

var APIRoutes = []Route{
	{
		URI:          "/store/site",
		Method:       http.MethodPost,
		Controller:   StoreWebsite,
		AuthRequired: false,
	},
	{
		URI:          "/all-sites",
		Method:       http.MethodGet,
		Controller:   RetriveAllWebsites,
		AuthRequired: false,
	},
	{
		URI:          "/get/site",
		Method:       http.MethodPost,
		Controller:   RetriveWebsiteByDomain,
		AuthRequired: false,
	},
}
