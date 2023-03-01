package main

import (
	"encoding/json"
	"net/http"
)

type student struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	BirthYear int    `json:"birth_year"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a new student object with sample data
		s := student{
			Name:      "John Doe",
			ID:        "12345",
			BirthYear: 2000,
		}

		// Convert the student object to JSON
		jsonData, err := json.Marshal(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON data to the response writer
		w.Write(jsonData)
	})

	// Start the HTTP server on localhost:8080
	http.ListenAndServe(":8080", nil)
}

/*
http://localhost:8080/
{"name":"John Doe","id":"12345","birth_year":2000}
*/
