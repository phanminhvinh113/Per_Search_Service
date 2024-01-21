package main

import (
	"fmt"
	"net/http"
	"os"
	"src/service/api"
	"src/service/config"
	"src/service/database"
)

func main() {

	config.LoadEnv()

	database.ConnectDB()

	r := api.Router()
	// Start the HTTP server
	http.Handle("/", r)
	http.ListenAndServe(os.Getenv("PORT"), nil)
	fmt.Println("Server is running on http://localhost:8666")
}
