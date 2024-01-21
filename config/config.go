package config

import (
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db                    *mongo.Database
	client                *mongo.Client
	productCollection     *mongo.Collection
	collection2           *mongo.Collection
	databaseName          = "Per-Ecommerce"
	productCollectionName = "Product"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
