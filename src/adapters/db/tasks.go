package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gsn_manager_service/src/adapters"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var TaskRepo *TaskRepository

func NewTaskRepository(client *mongo.Client, dbName, collectionName string) *TaskRepository {
	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	repository := &TaskRepository{
		client:     client,
		database:   database,
		collection: collection,
	}

	TaskRepo = repository

	return repository
}

// CreateTodo inserts a new todo into the collection
func (r *TaskRepository) CreateTodo(ctx context.Context, payload *CreateNewTask) (*Tasks, error) {
	now := time.Now()

	newTask := &Tasks{
		Title:     payload.Title,
		Timestamp: *payload.Timestamp,
		Completed: payload.Completed,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result, err := r.collection.InsertOne(ctx, newTask)
	if err != nil {
		adapters.Logger.Error().Msg(fmt.Sprintf("Error creating new task => %v", err))
		return nil, err
	}

	newTask.ID = result.InsertedID.(bson.ObjectID)
	return newTask, nil
}

// GetTodos retrieves all todos from the collection
func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]Tasks, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []Tasks
	for cursor.Next(ctx) {
		var task Tasks
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTodoByID retrieves a single todo by its ObjectID
func (r *TaskRepository) GetTaskById(ctx context.Context, id string) (*Tasks, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task Tasks
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTodo updates the title or completed status of a todo by its ID
func (r *TaskRepository) ModifyTask(ctx context.Context, id string, payload *UpdateTask) (*Tasks, error) {
	if payload.IsEmpty() {
		return nil, errors.New("payload can not be empty")
	}

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	set := bson.M{
		"updated_at": time.Now(),
	}

	if payload.Title != nil {
		set["title"] = *payload.Title
	}
	if payload.Completed != nil {
		set["completed"] = *payload.Completed
	}
	if payload.Timestamp != nil {
		set["timestamp"] = *payload.Timestamp
	}

	updateDoc := bson.M{"$set": set}

	filter := bson.M{"_id": objID}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedTask Tasks
	err = r.collection.FindOneAndUpdate(ctx, filter, updateDoc, opts).Decode(&updatedTask)
	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

// DeleteTodo removes a todo by its ID
func (r *TaskRepository) DeleteTask(ctx context.Context, id string) (bson.ObjectID, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return bson.NilObjectID, err
	}

	filter := bson.M{"_id": objID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return bson.NilObjectID, err
	}

	if result.DeletedCount == 0 {
		return bson.NilObjectID, mongo.ErrNoDocuments
	}

	return objID, nil
}
