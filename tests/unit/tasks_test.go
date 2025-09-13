package unit

import (
	"context"
	"errors"
	"testing"

	"github.com/gsn_manager_service/src/adapters/db"
	"github.com/gsn_manager_service/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestCreateTodo(t *testing.T) {
	ctx := context.Background()
	t.Run("Should create a new todo successfully in the db and return it.", func(t *testing.T) {
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
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.Title, payload.Title)
		assert.Equal(t, result.ID, expectedID)
		assert.Equal(t, result.Completed, payload.Completed)
		assert.False(t, result.UpdatedAt.IsZero())
		assert.False(t, result.Timestamp.IsZero())
		assert.False(t, result.CreatedAt.IsZero())
	})

	t.Run("Should return an error when there is a db connection failure.", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)

		payload := mocks.GetSampleCreateTaskPayload()
		expectedError := errors.New("Database connection failed")

		mockCollection.InsertOneFunc = func(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
			return nil, expectedError
		}

		// Act
		result, err := repo.CreateTodo(ctx, payload)

		// Assert
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), expectedError.Error())
	})
}

func TestGetAllTasks_Success(t *testing.T) {
	ctx := context.Background()
	t.Run("Should return a list of tasks successfully.", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)

		expectedTasks := mocks.GetMultipleTasks()

		mockCollection.FindFunc = func(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
			var docs []any
			for _, task := range expectedTasks {
				docs = append(docs, task)
			}

			cursor, err := mongo.NewCursorFromDocuments(docs, nil, nil)
			return cursor, err
		}

		// Act
		result, err := repo.GetAllTasks(ctx)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, len(result), len(expectedTasks))

		for i, task := range result {
			assert.Equal(t, task.Title, result[i].Title)
		}
	})

	t.Run("Should successfully return an empty slice when there is no tasks", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)
		emptyArr := make([]db.Tasks, 0)

		mockCollection.FindFunc = func(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
			var docs []any
			cursor, err := mongo.NewCursorFromDocuments(docs, nil, nil)
			return cursor, err
		}

		taskList, err := repo.GetAllTasks(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, taskList)
		assert.Equal(t, emptyArr, taskList)
		assert.Equal(t, len(taskList), 0)
	})

	t.Run("Should return an error when there is a db failure while retrieving the tasks", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)
		expectedDbError := errors.New("DB failure during")

		mockCollection.FindFunc = func(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
			return nil, expectedDbError
		}

		taskList, err := repo.GetAllTasks(ctx)

		assert.Nil(t, taskList)
		assert.NotNil(t, err)
		assert.Equal(t, expectedDbError.Error(), err.Error())
	})
}

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
