package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}
type Episode struct {
	ID          primitive.ObjectID `bson:"id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

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
	database := client.Database("quickstart")
	podcastsCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes")

	//Create a podcast
	mongopodcast := Podcast{
		Name:   "The mongo podcast",
		Author: "Yagnik P",
		Tags:   []string{"mongodb", "nosql"},
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, mongopodcast)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted result", insertResult.InsertedID)

	//Read all the podcast
	var podcasts []Podcast
	podcastCurser, err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	err = podcastCurser.All(ctx, &podcasts)
	if err != nil {
		panic(err)
	}
	fmt.Println(podcasts)

	var episodes []Episode
	episodeCurser, err := episodesCollection.Find(ctx, Episode{Duration: 25})
	if err != nil {
		panic(err)
	}
	err = episodeCurser.All(ctx, &episodes)
	if err != nil {
		panic(err)
	}
	fmt.Println("25 base filter response is", episodes)

}
