package adapters

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client

func ConnectToMongoDb(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("☠️ MongoDB connection failed:", err)
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("MongoDB ping error:", err)
		return nil, err
	}

	Client = client
	log.Println("📻 Connected to MongoDB!")
	return client, nil
}

func DisconnectMongo() {
	if Client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		log.Printf("⚠️ Error disconnecting MongoDB: %v", err)
	} else {
		log.Println("👋 Disconnected from MongoDB.")
	}
}
