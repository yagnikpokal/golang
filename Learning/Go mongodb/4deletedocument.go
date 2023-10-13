package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	//url := os.Getenv("mongodb+srv://yagnikpokal:yagnikpokal@cluster0.7baraeb.mongodb.net/?retryWrites=true&w=majority")
	opts := options.Client().ApplyURI("mongodb+srv://yagnikpokal:yagnikpokal@cluster0.7baraeb.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Use on linux command prompt export MONGODB_URI='mongodb+srv://yagnikpokal:yagnikpokal@cluster0.7baraeb.mongodb.net/?retryWrites=true&w=majority'
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("episodes")
	// Updating a single documents
	result, err := podcastsCollection.DeleteOne(
		ctx,
		bson.M{"duration": 25},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted the %v documents", result.DeletedCount)

	result, err = podcastsCollection.DeleteMany(
		ctx,
		bson.M{"duration": 32},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted the %v documents", result.DeletedCount)
}
