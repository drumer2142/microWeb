package api

import (
	"net/http"
)

func (srv *APIServer) LoadRoutes() []Route {

	var APIRoutes = []Route{
		{
			URI:          "/store/site",
			Method:       http.MethodPost,
			Controller:   srv.StoreWebsite,
			AuthRequired: false,
		},
		{
			URI:          "/all-sites",
			Method:       http.MethodGet,
			Controller:   srv.RetriveAllWebsites,
			AuthRequired: false,
		},
		{
			URI:          "/get/site",
			Method:       http.MethodPost,
			Controller:   srv.RetriveWebsiteByDomain,
			AuthRequired: false,
		},
	}

	return APIRoutes
}
