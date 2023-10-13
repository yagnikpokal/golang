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
	opts := options.Client().ApplyURI("ur;").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	// Define a filter to specify the document you want to retrieve
	filter := bson.D{{"title", "The Polyglot Developer Podcast"}}

	// Declare a variable to store the result
	var podcast bson.M

	// Use the FindOne method to retrieve one document
	err = podcastsCollection.FindOne(ctx, filter).Decode(&podcast)
	if err != nil {
		log.Fatal(err)
	}

	// Print the retrieved document
	fmt.Printf("ID: %s\nPodcast Title: %s\nAuthor: %s\n", podcast["_id"], podcast["title"], podcast["author"])

	// Read All documents
	filter = bson.D{}

	// Use the Find method to retrieve multiple documents
	cursor, err := episodesCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// Loop through the cursor to retrieve and print each document
	for cursor.Next(ctx) {
		var episode bson.M
		if err := cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Episode Title: %s\nDescription: %s\nDuration: %d\n", episode["title"], episode["description"], episode["duration"])
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Read all the Document where author is Nic Robby in documents
	// Define a filter to retrieve documents with "author" equal to "Nic Raboy"
	filter = bson.D{{"author", "Nic Raboy"}}

	// Use the Find method to retrieve all matching documents
	cursor, err = podcastsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// Loop through the cursor to retrieve and print each document
	for cursor.Next(ctx) {
		var podcast bson.M
		if err := cursor.Decode(&podcast); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Podcast Title: %s\nAuthor: %s\n", podcast["title"], podcast["author"])
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

}
