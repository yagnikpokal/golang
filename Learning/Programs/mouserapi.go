package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiEndpoint = "https://api.mouser.com/api/v1/search/keyword"
const apiKey = "5ad0a426-d6d8-4eda-acd2-7f9df0b07e50"

func main() {
	// Set up the search parameters
	params := map[string]string{
		"apiKey":  apiKey,
		"keyword": "solid state relay",
	}

	// Build the API URL with the search parameters
	url := apiEndpoint + "?" + buildQueryString(params)

	// Make the API request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the JSON response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract the data from the response
	products := response["products"].([]interface{})
	for _, product := range products {
		p := product.(map[string]interface{})
		partNumber := p["partNumber"].(string)
		manufacturer := p["manufacturer"].(string)
		fmt.Printf("%s - %s\n", partNumber, manufacturer)
	}
}

// buildQueryString creates a URL query string from a map of parameters
func buildQueryString(params map[string]string) string {
	var queryString string
	for key, value := range params {
		queryString += key + "=" + value + "&"
	}
	return queryString
}
