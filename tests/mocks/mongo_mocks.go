package mocks

import (
	"context"

	"github.com/gsn_manager_service/src/adapters/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestTaskRepository(client *mongo.Client, database *mongo.Database, collection db.CollectionInterface) *db.TaskRepository {
	return &db.TaskRepository{
		Client:     client,
		Database:   database,
		Collection: collection,
	}
}

// MongoCollection interface to abstract MongoDB collection operations
type MongoCollection interface {
	InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult
	FindOneAndUpdate(ctx context.Context, filter any, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult
	DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
}

// MockCollection implements MongoCollection for testing
type MockCollection struct {
	InsertOneFunc        func(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)
	FindFunc             func(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)
	FindOneFunc          func(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult
	FindOneAndUpdateFunc func(ctx context.Context, filter any, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult
	DeleteOneFunc        func(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error)
}

func (m *MockCollection) InsertOne(ctx context.Context, document any, opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error) {
	if m.InsertOneFunc != nil {
		return m.InsertOneFunc(ctx, document, opts...)
	}
	return nil, nil
}

func (m *MockCollection) Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	if m.FindFunc != nil {
		return m.FindFunc(ctx, filter, opts...)
	}
	return nil, nil
}

func (m *MockCollection) FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
	if m.FindOneFunc != nil {
		return m.FindOneFunc(ctx, filter, opts...)
	}
	return nil
}

func (m *MockCollection) FindOneAndUpdate(ctx context.Context, filter any, update any, opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult {
	if m.FindOneAndUpdateFunc != nil {
		return m.FindOneAndUpdateFunc(ctx, filter, update, opts...)
	}
	return nil
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter any, opts ...options.Lister[options.DeleteOneOptions]) (*mongo.DeleteResult, error) {
	if m.DeleteOneFunc != nil {
		return m.DeleteOneFunc(ctx, filter, opts...)
	}
	return nil, nil
}
