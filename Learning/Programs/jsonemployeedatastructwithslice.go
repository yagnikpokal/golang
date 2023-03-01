package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type employee struct {
	Title  string `json:"title"`
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Salary int    `json:"salary"`
}

type project struct {
	ProjectName string     `json:"projectName"`
	ProjectCode string     `json:"projectCode"`
	Employees   []employee `json:"employee"`
}

func main() {
	// Create the Winter project
	winterEmployees := []employee{
		{
			Title:  "Mr.",
			Name:   "A",
			ID:     5,
			Salary: 12345,
		},
		{
			Title:  "Mr.",
			Name:   "B",
			ID:     8,
			Salary: 54321,
		},
		{
			Title:  "Mr.",
			Name:   "C",
			ID:     9,
			Salary: 23456,
		},
	}

	winter := project{
		ProjectName: "Winter",
		ProjectCode: "O0123",
		Employees:   winterEmployees,
	}

	// Create the Summer project
	summerEmployees := []employee{
		{
			Title:  "Mr.",
			Name:   "D",
			ID:     6,
			Salary: 51234,
		},
		{
			Title:  "Mr.",
			Name:   "E",
			ID:     11,
			Salary: 7654321,
		},
	}

	summer := project{
		ProjectName: "Summer",
		ProjectCode: "P10406",
		Employees:   summerEmployees,
	}

	// Create a slice of projects
	projects := []project{winter, summer}

	// Create a handler function for the /data endpoint
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Convert the projects slice to JSON
		jsonData, err := json.Marshal(projects)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the content-type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON data to the response body
		w.Write(jsonData)
	})

	// Start the web server
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

/*

Go to http://localhost:8080/data

Responce will be like this

[{"projectName":"Winter","projectCode":"O0123","employee":[{"title":"Mr.","name":"A","id":5,"salary":12345},{"title":"Mr.","name":"B","id":8,"salary":54321},{"title":"Mr.","name":"C","id":9,"salary":23456}]},
{"projectName":"Summer","projectCode":"P10406","employee":[{"title":"Mr.","name":"D","id":6,"salary":51234},{"title":"Mr.","name":"E","id":11,"salary":7654321}]}]

*/
