package main

import "go.mongodb.org/mongo-driver/mongo/options"

func initializeDBOptions() *options.ClientOptions {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://yagnikpokal:yagnikpokal@cluster0.7baraeb.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	return opts
}