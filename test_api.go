package main

import (
	"fmt"
	"log"
	"net/http"
	"test/services/api"
)

func main() {
	// Initialize the database
	if err := api.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}

	// Register example endpoints
	http.HandleFunc("/api/examples", api.Handle(api.ExampleHandler))
	http.HandleFunc("/api/examples/create", api.Handle(api.CreateSampleHandler))

	// Start the server
	fmt.Println("API server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}