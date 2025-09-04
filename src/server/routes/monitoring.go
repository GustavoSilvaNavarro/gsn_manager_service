package routes

import (
	"log"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received health check request from %s", r.RemoteAddr)

	// Set the Content-Type header to plain text.
	w.Header().Set("Content-Type", "text/plain")

	// Write the "OK" response and a 200 OK status code.
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Error writing health check response: %v", err)
	}
}
