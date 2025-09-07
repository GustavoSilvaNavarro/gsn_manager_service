package routes

import (
	"fmt"
	"net/http"

	"github.com/gsn_manager_service/src/adapters"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		adapters.Logger.Error().Msg(fmt.Sprintf("Error writing health check response: %v", err))
	}
}
