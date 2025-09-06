package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gsn_manager_service/src/adapters/db"
	"github.com/gsn_manager_service/src/utils"
)

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload db.CreateNewTask
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Payload is missing -> Error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, "Payload is missing")
		return
	}

	fmt.Printf("%#v", payload)

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "Hello World"})
}
