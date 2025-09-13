package unit

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/gsn_manager_service/tests/mocks"
)

func TestCreateNewTask_Successfully(t *testing.T) {
	ctx := context.Background()
	mockCollection := &mocks.MockCollection{}
	repo := mocks.TestTaskRepository(nil, nil, mockCollection)

	payload := mocks.GetSampleCreateTaskPayload()
	expectedID := bson.NewObjectID()

	mockCollection.InsertOneFunc = func(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
		return &mongo.InsertOneResult{
			InsertedID: expectedID,
		}, nil
	}

	// Act
	result, err := repo.CreateTodo(ctx, payload)

	// Assert
	if err != nil {
		t.Errorf("Should not return an error but got %v", err)
	}

	if result == nil {
		t.Fatal("Should create and return a new task in the db.")
	}

	if result.Title != payload.Title {
		t.Errorf("Expected title %s, got %s", payload.Title, result.Title)
	}

	if result.ID != expectedID {
		t.Errorf("Expected ID %s, got %s", expectedID.Hex(), result.ID.Hex())
	}

	if result.Completed != payload.Completed {
		t.Errorf("Expecting completed property to be false but it received %v", result.Completed)
	}

	if result.CreatedAt.IsZero() {
		t.Errorf("CreatedAt field is zero, expected it to be set")
	}

	if result.UpdatedAt.IsZero() {
		t.Errorf("UpdatedAt field is zero, expected it to be set")
	}
}

// func TestCreateTodo_Error(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	mockCollection := &mocks.MockCollection{}
// 	repo := db.NewTaskRepositoryWithCollection(nil, nil, mockCollection)

// 	payload := fixtures.GetSampleCreateTaskPayload()
// 	expectedError := errors.New("database connection failed")

// 	mockCollection.InsertOneFunc = func(ctx context.Context, document interface{}, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
// 		return nil, expectedError
// 	}

// 	// Act
// 	result, err := repo.CreateTodo(ctx, payload)

// 	// Assert
// 	if err == nil {
// 		t.Error("Expected error, got nil")
// 	}

// 	if result != nil {
// 		t.Error("Expected result to be nil")
// 	}

// 	if err.Error() != expectedError.Error() {
// 		t.Errorf("Expected error %s, got %s", expectedError.Error(), err.Error())
// 	}
// }

// func TestGetAllTasks_Success(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	mockCollection := &mocks.MockCollection{}
// 	repo := db.NewTaskRepositoryWithCollection(nil, nil, mockCollection)

// 	expectedTasks := fixtures.GetMultipleTasks()

// 	mockCollection.FindFunc = func(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
// 		// Convert tasks to interface{} slice
// 		var docs []interface{}
// 		for _, task := range expectedTasks {
// 			docs = append(docs, task)
// 		}

// 		cursor, err := tests.NewMockCursor(docs)
// 		return cursor, err
// 	}

// 	// Act
// 	result, err := repo.GetAllTasks(ctx)

// 	// Assert
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}

// 	if len(result) != len(expectedTasks) {
// 		t.Errorf("Expected %d tasks, got %d", len(expectedTasks), len(result))
// 	}

// 	for i, task := range result {
// 		if task.Title != expectedTasks[i].Title {
// 			t.Errorf("Expected task title %s, got %s", expectedTasks[i].Title, task.Title)
// 		}
// 	}
// }

// func TestGetTaskById_Success(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	mockCollection := &mocks.MockCollection{}
// 	repo := db.NewTaskRepositoryWithCollection(nil, nil, mockCollection)

// 	expectedTask := fixtures.GetSampleTask()
// 	taskID := expectedTask.ID.Hex()

// 	mockCollection.FindOneFunc = func(ctx context.Context, filter interface{}, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
// 		return tests.NewMockSingleResult(expectedTask, nil)
// 	}

// 	// Act
// 	result, err := repo.GetTaskById(ctx, taskID)

// 	// Assert
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}

// 	if result == nil {
// 		t.Fatal("Expected result to be non-nil")
// 	}

// 	if result.Title != expectedTask.Title {
// 		t.Errorf("Expected title %s, got %s", expectedTask.Title, result.Title)
// 	}
// }

// func TestModifyTask_Success(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	mockCollection := &mocks.MockCollection{}
// 	repo := db.NewTaskRepositoryWithCollection(nil, nil, mockCollection)

// 	taskID := bson.NewObjectID().Hex()
// 	payload := fixtures.GetSampleUpdateTaskPayload()

// 	updatedTask := fixtures.GetSampleTask()
// 	updatedTask.Title = *payload.Title
// 	updatedTask.Completed = *payload.Completed

// 	mockCollection.FindOneAndUpdateFunc = func(ctx context.Context, filter interface{}, update interface{}, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult {
// 		return tests.NewMockSingleResult(updatedTask, nil)
// 	}

// 	// Act
// 	result, err := repo.ModifyTask(ctx, taskID, payload)

// 	// Assert
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}

// 	if result == nil {
// 		t.Fatal("Expected result to be non-nil")
// 	}

// 	if result.Title != *payload.Title {
// 		t.Errorf("Expected title %s, got %s", *payload.Title, result.Title)
// 	}
// }

// func TestModifyTask_EmptyPayload(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	mockCollection := &mocks.MockCollection{}
// 	repo := db.NewTaskRepositoryWithCollection(nil, nil, mockCollection)

// 	taskID := bson.NewObjectID().Hex()
// 	payload := &db.UpdateTask{} // Empty payload

// 	// Act
// 	result, err := repo.ModifyTask(ctx, taskID, payload)

// 	// Assert
// 	if err == nil {
// 		t.Error("Expected error for empty payload, got nil")
// 	}

// 	if result != nil {
// 		t.Error("Expected result to be nil")
// 	}

// 	expectedError := "payload can not be empty"
// 	if err.Error() != expectedError {
// 		t.Errorf("Expected error %s, got %s", expectedError, err.Error())
// 	}
// }

// func TestDeleteTask_Success(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	mockCollection := &mocks.MockCollection{}
// 	repo := db.NewTaskRepositoryWithCollection(nil, nil, mockCollection)

// 	taskID := bson.NewObjectID()

// 	mockCollection.DeleteOneFunc = func(ctx context.Context, filter interface{}, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
// 		return &mongo.DeleteResult{
// 			DeletedCount: 1,
// 		}, nil
// 	}

// 	// Act
// 	result, err := repo.DeleteTask(ctx, taskID.Hex())

// 	// Assert
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}

// 	if result != taskID {
// 		t.Errorf("Expected ID %s, got %s", taskID.Hex(), result.Hex())
// 	}
// }

// func TestDeleteTask_NotFound(t *testing.T) {
// 	// Arrange
// 	ctx := context.Background()
// 	mockCollection := &mocks.MockCollection{}
// 	repo := db.NewTaskRepositoryWithCollection(nil, nil, mockCollection)

// 	taskID := bson.NewObjectID()

// 	mockCollection.DeleteOneFunc = func(ctx context.Context, filter interface{}, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
// 		return &mongo.DeleteResult{
// 			DeletedCount: 0, // No documents deleted
// 		}, nil
// 	}

// 	// Act
// 	result, err := repo.DeleteTask(ctx, taskID.Hex())

// 	// Assert
// 	if err != mongo.ErrNoDocuments {
// 		t.Errorf("Expected ErrNoDocuments, got %v", err)
// 	}

// 	if result != bson.NilObjectID {
// 		t.Errorf("Expected NilObjectID, got %s", result.Hex())
// 	}
// }
