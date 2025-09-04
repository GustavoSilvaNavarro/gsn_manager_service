package routes

import "net/http"

// Setup all the routes here
func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", Healthz)
}
