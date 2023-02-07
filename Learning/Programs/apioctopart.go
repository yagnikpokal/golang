package main

import (
	"context"
	"fmt"
	//"github.com/marema31/octopart"
)

const apiKey = "YOUR_API_KEY_HERE"

func main() {
	// Set up the client
	client := octopart.NewClient(nil, apiKey)

	// Set up the request parameters
	params := &octopart.PartMatchParams{
		Queries: []octopart.PartMatchQuery{
			{
				Manufacturer: "Texas Instruments",
				MPN:          "SN74LS00N",
			},
		},
		Include: []string{"specs"},
	}

	// Send the request and get the response
	res, _, err := client.PartMatch.Get(context.Background(), params)
	if err != nil {
		panic(err)
	}

	// Print the response body
	fmt.Println(res)
}
