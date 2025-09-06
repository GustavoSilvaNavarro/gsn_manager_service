package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
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
		log.Printf("Invalid payload -> Error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	// Validate required fields
	if err := utils.Validate.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)
		msg := ""
		for _, e := range errs {
			if utf8.RuneCountInString(msg) > 0 {
				msg += " | "
			}
			msg += fmt.Sprintf("%s is %s", e.Field(), e.Tag())
		}
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	newTask, err := db.TaskRepo.CreateTodo(r.Context(), &payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create task")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, newTask)
}
