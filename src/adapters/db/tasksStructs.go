package db

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// ? DB Model for Task
type Tasks struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Completed bool          `bson:"completed" json:"completed"`
	Timestamp time.Time     `bson:"timestamp" json:"timestamp"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

type TaskRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// ? Struct for new task
type CreateNewTask struct {
	Title     string    `json:"title" validate:"required"`
	Timestamp time.Time `json:"timestamp" validate:"required"`
	Completed bool      `json:"completed"`
}

// ? Struct to update task
type UpdateTask struct {
	Title     *string    `json:"title,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Completed *bool      `json:"completed,omitempty"`
}
