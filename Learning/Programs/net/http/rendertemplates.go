// Templates are located in the net/http/templates
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const Port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTempate(w, "home.page.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	RenderTempate(w, "about.page.html")
}
func RenderTempate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in parsing the templates", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(Port, nil)
}

/*
http://localhost:8080/
This is the home page

http://localhost:8080/about
This is the about page
Contact us on the below
*/
