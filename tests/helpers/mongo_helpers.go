package helpers

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// MockCursor implements mongo.Cursor for testing
type MockCursor struct {
	data    []bson.Raw
	current int
	closed  bool
}

func NewMockCursor(documents []any) (*MockCursor, error) {
	var rawDocs []bson.Raw

	for _, doc := range documents {
		rawDoc, err := bson.Marshal(doc)
		if err != nil {
			return nil, err
		}
		rawDocs = append(rawDocs, rawDoc)
	}

	return &MockCursor{
		data:    rawDocs,
		current: -1,
		closed:  false,
	}, nil
}

func (c *MockCursor) Close(ctx context.Context) error {
	c.closed = true
	return nil
}

func (c *MockCursor) Next(ctx context.Context) bool {
	if c.closed {
		return false
	}

	c.current++
	return c.current < len(c.data)
}

func (c *MockCursor) Decode(val any) error {
	if c.current < 0 || c.current >= len(c.data) {
		return mongo.ErrNoDocuments
	}

	return bson.Unmarshal(c.data[c.current], val)
}

// MockSingleResult implements mongo.SingleResult for testing
type MockSingleResult struct {
	data []byte
	err  error
}

func NewMockSingleResult(document any, err error) *MockSingleResult {
	if err != nil {
		return &MockSingleResult{err: err}
	}

	data, marshalErr := bson.Marshal(document)
	if marshalErr != nil {
		return &MockSingleResult{err: marshalErr}
	}

	return &MockSingleResult{data: data}
}

func (r *MockSingleResult) Decode(v any) error {
	if r.err != nil {
		return r.err
	}

	return bson.Unmarshal(r.data, v)
}
