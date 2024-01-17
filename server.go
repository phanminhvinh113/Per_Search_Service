package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
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
	// Specify the database and collection you want to use
	database := client.Database("Per-Ecommerce")
	collection = database.Collection("Product")

	//
	fmt.Println("Connected to MongoDB!")
}

type Product struct {
	ID       string `bson:"_id" json:"_id"`
	Name     string `bson:"name" json:"name"`
	Price    int    `bson:"price" json:"price"`
	Quantity int    `bson:"quantity" json:"quantity"`
}

// Example search function
func searchProductsByName(keyword string) ([]Product, error) {

	// MongoDB query to find products by name
	filter := bson.M{
		"$text": bson.M{
			"$search": keyword,
		},
	}
	options := options.Find().SetLimit(20)
	// Execute the query
	cursor, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}

	var products []Product
	// Iterate through the results
	for cursor.Next(context.TODO()) {
		var product Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
func searchHandler(w http.ResponseWriter, r *http.Request) {

	// Extract search query from URL
	keyword := r.URL.Query().Get("keyword")

	// Call the search function
	results, err := searchProductsByName(keyword)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return results as JSON
	json.NewEncoder(w).Encode(results)
}
func main() {

	config.loadEnv()
	database.connectDB()

	r := mux.NewRouter()

	// Define a route that will handle concurrent requests
	r.HandleFunc("/concurrent", ConcurrentHandler).Methods("GET")
	r.HandleFunc("/search", searchHandler).Methods("GET")
	// Start the HTTP server
	http.Handle("/", r)
	fmt.Println("Server is running on http://localhost:8666")
	http.ListenAndServe(os.Getenv("PORT"), nil)
}

// ConcurrentHandler handles concurrent requests using Goroutines
func ConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	// Use a WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Number of concurrent Goroutines to spawn
	numRequests := 10

	// Increment the WaitGroup counter
	wg.Add(numRequests)

	// Handle each request concurrently
	for i := 0; i < numRequests; i++ {
		go func(index int) {
			// Decrement the WaitGroup counter when the Goroutine completes
			defer wg.Done()

			// Simulate some processing time
			time.Sleep(time.Second)

			// Respond to the client
			response := fmt.Sprintf("Response from Goroutine %d", index)
			fmt.Fprintln(w, response)
		}(i)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
}
