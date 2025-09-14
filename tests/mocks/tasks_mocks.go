package mocks

import (
	"time"

	"github.com/gsn_manager_service/src/adapters/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetSampleTask(title string) *db.Tasks {
	now := time.Now()

	return &db.Tasks{
		ID:        bson.NewObjectID(),
		Title:     title,
		Completed: false,
		Timestamp: now,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func GetSampleCreateTaskPayload() *db.CreateNewTask {
	timestamp := time.Now()

	return &db.CreateNewTask{
		Title:     "New Task",
		Timestamp: &timestamp,
		Completed: false,
	}
}

func GetSampleUpdateTaskPayload() *db.UpdateTask {
	title := "Updated Task"
	completed := true
	timestamp := time.Now()

	return &db.UpdateTask{
		Title:     &title,
		Completed: &completed,
		Timestamp: &timestamp,
	}
}

func GetMultipleTasks() []db.Tasks {
	return []db.Tasks{*GetSampleTask("Task 1"), *GetSampleTask("Task 2")}
}
