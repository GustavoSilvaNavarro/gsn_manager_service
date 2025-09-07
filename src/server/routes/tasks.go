package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gsn_manager_service/src/adapters"
	"github.com/gsn_manager_service/src/adapters/db"
	"github.com/gsn_manager_service/src/utils"
)

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var payload db.CreateNewTask
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		adapters.Logger.Error().Msg(fmt.Sprintf("Invalid payload -> Error: %v", err))
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

func RetrieveAllTasks(w http.ResponseWriter, r *http.Request) {
	allTasks, err := db.TaskRepo.GetAllTasks(r.Context())

	if err != nil {
		adapters.Logger.Error().Msg(fmt.Sprintf("Error calling the tasks => %v", err))
		utils.WriteError(w, http.StatusInternalServerError, "Failed to retrieve all tasks")
		return
	}

	utils.WriteJSON(w, http.StatusOK, allTasks)
}

func GetSingleTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, err := db.TaskRepo.GetTaskById(r.Context(), id)
	if err != nil {
		msg := fmt.Sprintf("Failed to find a task with ID: %s | Error => %v", id, err.Error())
		adapters.Logger.Error().Msg(msg)
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var payload db.UpdateTask
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		adapters.Logger.Error().Msg(fmt.Sprintf("Invalid payload -> Error: %v", err.Error()))
		utils.WriteError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)
		msg := ""
		for _, e := range errs {
			if utf8.RuneCountInString(msg) > 0 {
				msg += " | "
			}
			msg += fmt.Sprintf("%s is %s", e.Field(), e.Tag())
		}
		message := fmt.Sprintf("Failed to update task with ID: %s. Error => ", id)
		utils.WriteError(w, http.StatusBadRequest, message+msg)
		return
	}

	updatedTask, err := db.TaskRepo.ModifyTask(r.Context(), id, &payload)
	if err != nil {
		msg := fmt.Sprintf("Failed to update task with ID: %s | Error => %v", id, err.Error())
		adapters.Logger.Error().Msg(msg)
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &updatedTask)
}

func RemoveTaskById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	taskId, err := db.TaskRepo.DeleteTask(r.Context(), id)
	if err != nil {
		msg := fmt.Sprintf("Failed to remove task with ID: %s | Error => %v", id, err.Error())
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Successfully removed task with ID: %s", taskId.Hex()),
	})
}
