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
	Ctx                       = context.Background()
	CustomersCollection       *mongo.Collection
	CyclesCollection          *mongo.Collection
	NotificationsCollection   *mongo.Collection
	PhishesCollection         *mongo.Collection
	SendingProfilesCollection *mongo.Collection
	SubscriptionsCollection   *mongo.Collection
	TargetsCollection         *mongo.Collection
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
	CustomersCollection = db.Collection("customer")
	CyclesCollection = db.Collection("cycle")
	NotificationsCollection = db.Collection("notification")
	PhishesCollection = db.Collection("template")
	SendingProfilesCollection = db.Collection("sending_profile")
	SubscriptionsCollection = db.Collection("subscription")
	TargetsCollection = db.Collection("target")
}
