package config

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	db                    *mongo.Database
	client                *mongo.Client
	productCollection     *mongo.Collection
	collection2           *mongo.Collection
	databaseName          = "Per-Ecommerce"
	productCollectionName = "Product"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
