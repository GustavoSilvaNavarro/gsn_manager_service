package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Add this interface to your db package
type CollectionInterface interface {
	InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult
	FindOneAndUpdate(ctx context.Context, filter any, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult
	DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
}

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
	collection CollectionInterface
}

// ? Struct for new task
type CreateNewTask struct {
	Title     string     `json:"title" validate:"required,min=3,max=100"`
	Timestamp *time.Time `json:"timestamp" validate:"required"`
	Completed bool       `json:"completed"`
}

// ? Struct to update task
type UpdateTask struct {
	Title     *string    `json:"title,omitempty" validate:"omitempty,min=3,max=100"`
	Timestamp *time.Time `json:"timestamp,omitempty" validate:"omitempty"`
	Completed *bool      `json:"completed,omitempty" validate:"omitempty"`
}

func (u *UpdateTask) IsEmpty() bool {
	return u.Title == nil && u.Timestamp == nil && u.Completed == nil
}
