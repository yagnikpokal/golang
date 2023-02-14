package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}
func Add(x int, y int) int {
	return x + y
}
func About(w http.ResponseWriter, r *http.Request) {
	sum := Add(1, 3)
	fmt.Fprintf(w, fmt.Sprintf("3 + 1 is %d", sum))
}

const PortNumber = ":8080"

func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)
	fmt.Println("starting application on " + fmt.Sprintf(PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}

/*
http://localhost:8080/about
3 + 1 is 4

http://localhost:8080/home
This is the home page

*/
