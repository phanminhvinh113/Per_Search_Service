package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client                *mongo.Client
	productCollection     *mongo.Collection
	databaseName          = "Per-Ecommerce"
	productCollectionName = "Product"
)

func connectDB() {

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// Specify the database
	database := client.Database(databaseName)
	productCollection = database.Collection(productCollectionName)
	//
	fmt.Println("Connected to MongoDB!")
}
