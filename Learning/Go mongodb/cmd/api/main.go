package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

const port = 8080

type Connection struct {
	Movie *mongo.Collection
	User  *mongo.Collection
}
type application struct {
	DSN          string
	Domain       string
	DB           Connection
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	APIKey       string
}

func main() {
	// Set application config
	var app application

	// Initialize MongoDB options
	opts := initializeDBOptions()

	// Create a new client and connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Set the MongoDB database and collections
	app.DB.Movie = client.Database("movieapp").Collection("movie")
	app.DB.User = client.Database("movieapp").Collection("user")

	log.Println("Starting application on port", port)

	// Start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
