package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const ApiSector = "healthcare" // "healthcare", "veterinary", "insurance" ...

// Supported content types and accepts
var supportedContentTypes = []string{"json", "jwt", "jwm", "didcomm"}
var supportedAccepts = append(supportedContentTypes) // Add any other accept types if required.

type Route struct {
	Method  string
	Url     string
	Handler http.HandlerFunc
}

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType(supportedContentTypes...))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Handle("/metrics", promhttp.Handler())

	routes := []Route{
		{
			Method:  "POST",
			Url:     "/{vaultAlternateName}/cds-{territory}/{apiVersion}/{sector}/{specification}/connection/{connectionId}",
			Handler: HandlePostPrimaryDocToStore, // Changed to HandlePostPrimaryDocToVaultEDV
		},
	}

	for _, route := range routes {
		switch route.Method {
		case "POST":
			r.Post(route.Url, route.Handler)
		case "GET":
			r.Get(route.Url, route.Handler)
			// Add more HTTP methods as needed...
		}
	}

	return r
}

/*
func unsupportedSpecHandler(w http.ResponseWriter, r *http.Request) {
	var status = "400"
	var code = "400"
	var title = "Unsupported Specification"
	var detail = "The requested specification is not supported."

	primaryDoc := &didcommUtils.PrimaryDocument{
		Data: []didcommUtils.ResourceObject{},
		Errors: &[]didcommUtils.ErrorObject{
			didcommUtils.ErrorObject{
				Status: &status,
				Code:   &code,
				Title:  &title,
				Detail: &detail,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(primaryDoc)
}
*/
