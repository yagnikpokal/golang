package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	id, _ := primitive.ObjectIDFromHex("651ffda17cb4b42ce3d0176b")
	// Updating a single documents
	result, err := podcastsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"author", "Nicolas Robby"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated the %v documents", result.ModifiedCount)

	// Updating mutiple document
	result, err = podcastsCollection.UpdateMany(
		ctx,
		bson.M{"title": "The polygot developer podcast"},
		bson.D{
			{"$set", bson.D{{"author", "Nic Robby"}}},
		},
	)
	fmt.Printf("Updated the %v documents", result.ModifiedCount)

	result, err = podcastsCollection.ReplaceOne(
		ctx,
		bson.M{"author": "Nic robby"},
		bson.M{
			"title":  "The Nic Robby Show",
			"author": "Nicolas robby",
		},
	)
	fmt.Printf("Updated the %v documents", result.ModifiedCount)
}
