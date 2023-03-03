package main

import (
	"html/template"
	"log"
	"net/http"
)

type HomePageData struct {
	Title string
}

func main() {
	templates := template.Must(template.ParseFiles("index.html"))

	homePageData := HomePageData{
		Title: "Welcome to My Website!",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Get the value of the input field from the request body
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			inputValue := r.Form.Get("input-field")

			// Log the input value to the console
			log.Printf("Input value: %s", inputValue)
		}

		err := templates.ExecuteTemplate(w, "index.html", homePageData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
