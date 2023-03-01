package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RequestBody struct {
	SearchByPartRequest SearchByPartRequest `json:"SearchByPartRequest"`
}

type SearchByPartRequest struct {
	MouserPartNumber string `json:"mouserPartNumber"`
}

func (app *Config) routes() http.Handler {
	// Create the router
	mux := chi.NewRouter()

	// Set up the middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)
	// Define application routes
	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate", app.RegisterPage)
	mux.Mount("/members", app.authRouter())

	mux.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		fmt.Println(query)
		fmt.Fprint(w, "Search query: "+query)
		querytomouser(query)

		/*switch query {
		case query:
			fmt.Println("Some string is come")
			querytomouser(query)
			//mux.Get("/", handleIndex)

		}*/
	})

	mux.Get("/parts", func(w http.ResponseWriter, r *http.Request) {

		parts := getdata()
		json.NewEncoder(w).Encode(parts)
	})

	/*
		This function will send an email when we go to link localhost:8080/test-email
		mux.Get("/test-email", func(w http.ResponseWriter, r *http.Request) {
			m := Mail{
				Domain:      "localhost",
				Host:        "localhost",
				Port:        1025,
				Encryption:  "none",
				FromAddress: "info@mycompany.com",
				FromName:    "info",
				ErrorChan:   make(chan error),
			}
			msg := Message{
				To:      "me@here.com",
				Subject: "Test email",
				Data:    "Hello, World.",
			}
			m.sendMail(msg, make(chan error))
		})
	*/
	return mux
}

func (app *Config) authRouter() http.Handler {
	mux := chi.NewRouter()
	mux.Use(app.Auth)
	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)
	return mux
}

// This will query to mouser.com for the selected part number
func querytomouser(query string) {
	url := "https://api.mouser.com/api/v1/search/partnumber?apiKey=5ad0a426-d6d8-4eda-acd2-7f9df0b07e50"
	requestBody := RequestBody{SearchByPartRequest: SearchByPartRequest{MouserPartNumber: query}}
	jsonValue, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	printSelectedAttributesLineByLine(string(body))
	//sendAttributesToFrontend(string(body))

}

/*
	func handleIndex(w http.ResponseWriter, r *http.Request) {
		url := "https://api.mouser.com/api/v1/search/partnumber?apiKey=5ad0a426-d6d8-4eda-acd2-7f9df0b07e50"
		requestBody := RequestBody{SearchByPartRequest: SearchByPartRequest{MouserPartNumber: "BQ25892"}}
		jsonValue, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
		tmpl, _ := template.ParseFiles("index.gohtml")
		tmpl.Execute(w, string(body))
		// printAttributeLineByLine(string(body))
		printSelectedAttributesLineByLine(string(body))

}
*/

/*
	func handleIndex(w http.ResponseWriter, r *http.Request) {
		url := "https://api.mouser.com/api/v1/search/partnumber?apiKey=5ad0a"
		requestBody := RequestBody{SearchByPartRequest: SearchByPartRequest{MouserPartNumber: "BQ25892"}}
		jsonValue, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		printSelectedAttributesLineByLine(string(body))
	}
*/

func printSelectedAttributesLineByLine(body string) {
	var jsonBody map[string]interface{}
	json.Unmarshal([]byte(body), &jsonBody)
	searchResults := jsonBody["SearchResults"].(map[string]interface{})["Parts"].([]interface{})
	for _, v := range searchResults {
		part := v.(map[string]interface{})
		fmt.Println("Availability:", part["Availability"])
		fmt.Println("Description:", part["Description"])
		fmt.Println("ManufacturerPartNumber:", part["ManufacturerPartNumber"])
		fmt.Println("MouserPartNumber:", part["MouserPartNumber"])
		fmt.Println("Min:", part["Min"])

		priceBreaks := part["PriceBreaks"].([]interface{})
		for _, price := range priceBreaks {
			fmt.Println("Quantity:", price.(map[string]interface{})["Quantity"])
			fmt.Println("Price:", price.(map[string]interface{})["Price"])
		}
	}
}

type Part struct {
	Quantity     int     `json:"Quantity"`
	Availability string  `json:"Availability"`
	Description  string  `json:"Description"`
	Price        float64 `json:"Price"`
}

var newPart Part

func getdata() []Part {
	parts := []Part{
		//{Quantity: 20, Availability: "In Stock", Description: "Part 1", Price: 5.99},
		//{Quantity: 20, Availability: "Out of Stock", Description: "Part 2", Price: 10.99},
		//{Quantity: 30, Availability: "In Stock", Description: "Part 3", Price: 15.99},
	}
	newPart = Part{Quantity: 40, Availability: "In Stock", Description: "Part 4", Price: 20.99}

	parts = append(parts, newPart)

	return parts
}

/*
func sendAttributesToFrontend(body string) {
	var jsonBody map[string]interface{}
	json.Unmarshal([]byte(body), &jsonBody)
	searchResults := jsonBody["SearchResults"].(map[string]interface{})["Parts"].([]interface{})
	attributes := make([]map[string]interface{}, len(searchResults))
	for i, v := range searchResults {
		part := v.(map[string]interface{})
		attribute := make(map[string]interface{})
		attribute["Availability"] = part["Availability"]
		attribute["Description"] = part["Description"]
		attribute["ManufacturerPartNumber"] = part["ManufacturerPartNumber"]
		attribute["MouserPartNumber"] = part["MouserPartNumber"]
		attribute["Min"] = part["Min"]
		priceBreaks := part["PriceBreaks"].([]interface{})
		for _, price := range priceBreaks {
			attribute["Quantity"] = price.(map[string]interface{})["Quantity"]
			attribute["Price"] = price.(map[string]interface{})["Price"]
		}
		attributes[i] = attribute
	}
	// Use a package like "net/http" or "github.com/gin-gonic/gin" to send the attributes to the frontend here
}*/
/*
func getSelectedAttributes(body string) (map[string]interface{}, error) {
	var jsonBody map[string]interface{}
	err := json.Unmarshal([]byte(body), &jsonBody)
	if err != nil {
		return nil, err
	}

	searchResults := jsonBody["SearchResults"].(map[string]interface{})["Parts"].([]interface{})
	var parts []map[string]interface{}
	for _, v := range searchResults {
		part := v.(map[string]interface{})
		var priceBreaks []map[string]interface{}
		for _, price := range part["PriceBreaks"].([]interface{}) {
			priceBreaks = append(priceBreaks, price.(map[string]interface{}))
		}
		partData := map[string]interface{}{
			"Availability":           part["Availability"],
			"Description":            part["Description"],
			"ManufacturerPartNumber": part["ManufacturerPartNumber"],
			"MouserPartNumber":       part["MouserPartNumber"],
			"Min":                    part["Min"],
			"PriceBreaks":            priceBreaks,
		}
		parts = append(parts, partData)
		fmt.Println(parts)
		fmt.Println("This is part")
	}

	return map[string]interface{}{
		"Parts": parts,
	}, nil
}*/

/*
// Print all the attributes
func printAttributeLineByLine(body string) {
	var jsonBody interface{}
	json.Unmarshal([]byte(body), &jsonBody)
	searchResults := jsonBody.(map[string]interface{})["SearchResults"].(map[string]interface{})["Parts"].([]interface{})
	for _, v := range searchResults {
		productAttributes := v.(map[string]interface{})["ProductAttributes"].([]interface{})
		for _, pa := range productAttributes {
			attribute := pa.(map[string]interface{})
			fmt.Println("attributeName:", attribute["attributeName"])
			fmt.Println("attributeValue:", attribute["attributeValue"])
		}
	}
}*/
