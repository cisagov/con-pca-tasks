package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Ctx                     = context.Background()
	NotificationsCollection *mongo.Collection
	PhishesCollection       *mongo.Collection
	SubscriptionsCollection *mongo.Collection
	CyclesCollection        *mongo.Collection
)

func connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

// InitDB initializes database connection and sets the collections
func InitDB() {
	client, err := connect()
	if err != nil {
		panic(err)
	}
	db := client.Database("pca")

	// Set the collections
	CyclesCollection = db.Collection("cycle")
	NotificationsCollection = db.Collection("notification")
	PhishesCollection = db.Collection("template")
	SubscriptionsCollection = db.Collection("subscription")
}
