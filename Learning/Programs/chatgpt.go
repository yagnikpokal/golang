package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Set up the request URL and parameters
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"
	apiKey := "Your API"
	prompt := "Pin to pin compatible part for the bq25892"

	// Create the request body as a JSON object
	requestBody := `{
		"prompt": "` + prompt + `",
		"max_tokens": 250,
		"temperature": 0.5
	}`

	// Create the HTTP request with the necessary headers and body
	req, err := http.NewRequest("POST", url, ioutil.NopCloser(
		bytes.NewBuffer([]byte(requestBody))))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the HTTP request and retrieve the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body and print it to the command line
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println(string(respBody))
	fmt.Println()

	printText(string(respBody))
}

func printText(jsonStr string) error {
	var resp map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		return err
	}

	choices := resp["choices"].([]interface{})
	choice := choices[0].(map[string]interface{})
	text := choice["text"].(string)

	fmt.Println(text)

	return nil
}

/*
func printGeneratedText(jsonStr string) error {
	// Define a struct to represent the JSON response
	type jsonResponse struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	// Unmarshal the JSON response into the struct
	var response jsonResponse
	if err := json.Unmarshal([]byte(jsonStr), &response); err != nil {
		return err
	}

	// Print the generated text
	//fmt.Println(response.Choices[0].Text)
	//printText(response.Choices[0].Text)
	var resp Response
	if err := json.Unmarshal([]byte(response.Choices[0].Text), &resp); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(resp.Choices[0].Text)
	return nil

}*/

/*
func printText(jsonStr string) error {
	var resp map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		return err
	}

	choices := resp["choices"].([]interface{})
	choice := choices[0].(map[string]interface{})
	text := choice["text"].(string)

	fmt.Println(text)

	return nil

}
*/
