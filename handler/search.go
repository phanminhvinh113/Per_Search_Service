package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"src/service/database"
	"src/service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchProductsByName(keyword string) ([]utils.Product, error) {

	// MongoDB query to find products by name
	filter := bson.M{
		"$text": bson.M{
			"$search": keyword,
		},
	}
	options := options.Find().SetLimit(20)

	// Execute the query
	cursor, err := database.ProductCollection.Find(context.TODO(), filter, options)

	if err != nil {
		return nil, err
	}

	var products []utils.Product
	// Iterate through the results
	for cursor.Next(context.TODO()) {
		var product utils.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Extract search query from URL
	keyword := r.URL.Query().Get("keyword")

	// Call the search function
	results, err := SearchProductsByName(keyword)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return results as JSON
	json.NewEncoder(w).Encode(results)
}
