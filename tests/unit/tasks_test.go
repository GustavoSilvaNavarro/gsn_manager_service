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
		assert.NoError(t, err, "Should not return any errors")
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

func TestGetAllTasks(t *testing.T) {
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

func TestGetTaskById(t *testing.T) {
	ctx := context.Background()
	t.Run("Should return a task when valid task_id is passed.", func(t *testing.T) {
		// Arrange
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)

		expectedTask := mocks.GetSampleTask("Task 1")
		taskID := expectedTask.ID.Hex()

		mockCollection.FindOneFunc = func(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
			return mongo.NewSingleResultFromDocument(expectedTask, nil, bson.NewRegistry())
		}

		// Act
		result, err := repo.GetTaskById(ctx, taskID)

		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.Title, expectedTask.Title)
		assert.Equal(t, result.ID.Hex(), taskID)
		assert.Equal(t, result.Completed, expectedTask.Completed)
		assert.False(t, result.UpdatedAt.IsZero())
		assert.False(t, result.Timestamp.IsZero())
		assert.False(t, result.CreatedAt.IsZero())
	})
}

func TestModifyTask(t *testing.T) {
	ctx := context.Background()
	t.Run("Should return and updated task successfully.", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)

		taskID := bson.NewObjectID().Hex()
		payload := mocks.GetSampleUpdateTaskPayload()

		updatedTask := mocks.GetSampleTask("Task 1")
		updatedTask.Title = *payload.Title
		updatedTask.Completed = *payload.Completed
		updatedTask.Timestamp = *payload.Timestamp

		mockCollection.FindOneAndUpdateFunc = func(ctx context.Context, filter any, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult {
			return mongo.NewSingleResultFromDocument(updatedTask, nil, bson.NewRegistry())
		}

		// Act
		taskUpdated, err := repo.ModifyTask(ctx, taskID, payload)

		assert.Nil(t, err)
		assert.NotNil(t, taskUpdated)
		assert.Equal(t, taskUpdated.Title, updatedTask.Title)
		assert.Equal(t, taskUpdated.Completed, updatedTask.Completed)
		assert.Equal(t, taskUpdated.ID, updatedTask.ID)
		assert.False(t, taskUpdated.Timestamp.IsZero())
		assert.False(t, taskUpdated.CreatedAt.IsZero())
		assert.False(t, taskUpdated.UpdatedAt.IsZero())
	})

	t.Run("Should return an error when payload is empty", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)
		taskId := bson.NewObjectID()
		payload := &db.UpdateTask{}

		updatedTask, err := repo.ModifyTask(ctx, taskId.Hex(), payload)

		assert.NotNil(t, err)
		assert.Nil(t, updatedTask)
		assert.Equal(t, err.Error(), "payload can not be empty")
	})
}

func TestDeleteTask(t *testing.T) {
	ctx := context.Background()
	t.Run("Should successfully delete a task and return their ID", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)
		taskID := bson.NewObjectID()

		mockCollection.DeleteOneFunc = func(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
			return &mongo.DeleteResult{
				DeletedCount: 1,
			}, nil
		}

		deletedId, err := repo.DeleteTask(ctx, taskID.Hex())

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, deletedId, taskID)
	})

	t.Run("Should return an error when delete process has failed", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)
		taskID := bson.NewObjectID()

		mockCollection.DeleteOneFunc = func(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
			return nil, errors.New("Error deleting task")
		}

		deletedId, err := repo.DeleteTask(ctx, taskID.Hex())

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Error deleting task")
		assert.Equal(t, deletedId, bson.NilObjectID)
	})

	t.Run("Should return an error when no documents to get deleted were found", func(t *testing.T) {
		mockCollection := &mocks.MockCollection{}
		repo := mocks.TestTaskRepository(nil, nil, mockCollection)
		taskID := bson.NewObjectID()

		mockCollection.DeleteOneFunc = func(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
			return &mongo.DeleteResult{DeletedCount: 0}, nil
		}

		deletedId, err := repo.DeleteTask(ctx, taskID.Hex())

		// Assert
		assert.Equal(t, err, mongo.ErrNoDocuments)
		assert.Equal(t, deletedId, bson.NilObjectID)
		assert.Equal(t, err.Error(), "mongo: no documents in result")
	})
}
