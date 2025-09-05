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

func ConnectToMongoDb(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("☠️ MongoDB connection failed:", err)
		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("MongoDB ping error:", err)
		return err
	}

	Client = client
	log.Println("📻 Connected to MongoDB!")
	return nil
}
