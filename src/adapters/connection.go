package adapters

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	Client *mongo.Client
	onceDb sync.Once
)

func ConnectToMongoDb(url string) (*mongo.Client, error) {
	var err error
	onceDb.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		Client, err = mongo.Connect(options.Client().ApplyURI(url))
		if err != nil {
			Logger.Fatal().Msg(fmt.Sprintf("☠️ MongoDB connection failed: %v", err))
			return
		}

		if err = Client.Ping(ctx, readpref.Primary()); err != nil {
			Logger.Fatal().Msg(fmt.Sprintf("MongoDB ping error: => %v", err))
			return
		}

		Logger.Info().Msg("📻 Connected to MongoDB!")
	})

	return Client, err
}

func DisconnectMongo(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		Logger.Warn().Msg(fmt.Sprintf("⚠️ Error disconnecting MongoDB: %v", err))
	} else {
		Logger.Info().Msg("👋 Disconnected from MongoDB.")
	}
}
