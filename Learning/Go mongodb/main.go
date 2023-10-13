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
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(podcastResult.InsertedID)
	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "GraphQL for API Development"},
			{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Progressive Web Application Development"},
			{"description", "Learn about PWA development with Tara Manicsic."},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))
}
