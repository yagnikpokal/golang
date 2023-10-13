package main

import "go.mongodb.org/mongo-driver/mongo/options"

func initializeDBOptions() *options.ClientOptions {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("url").SetServerAPIOptions(serverAPI)
	return opts
}
