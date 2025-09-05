package db

import (
	"context"
	"time"

	"github.com/gsn_manager_service/src/adapters"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/bson/primitive"
)

// Todo represents a single todo item
type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Completed bool               `bson:"completed" json:"completed"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}

// CreateTodo inserts a new todo into the collection
func CreateTodo(todo Todo) (*primitive.ObjectID, error) {
	collection := adapters.GetCollection("mydatabase", "todos")
	todo.Timestamp = time.Now()

	result, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)
	return &id, nil
}

// GetTodos retrieves all todos from the collection
func GetTodos() ([]Todo, error) {
	collection := adapters.GetCollection("mydatabase", "todos")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var todos []Todo
	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// GetTodoByID retrieves a single todo by its ObjectID
func GetTodoByID(id string) (*Todo, error) {
	collection := adapters.GetCollection("mydatabase", "todos")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var todo Todo
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// UpdateTodo updates the title or completed status of a todo by its ID
func UpdateTodo(id string, updatedData map[string]interface{}) error {
	collection := adapters.GetCollection("mydatabase", "todos")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": updatedData}
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	return err
}

// DeleteTodo removes a todo by its ID
func DeleteTodo(id string) error {
	collection := adapters.GetCollection("mydatabase", "todos")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}
