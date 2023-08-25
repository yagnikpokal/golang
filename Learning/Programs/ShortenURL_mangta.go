package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
)

var (
	urlMap = make(map[string]string)
)

func hasURL(url string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(url))
	return h.Sum32()
}

func generateShortenURL(longURL string) string {
	hash := hasURL(longURL)
	shortURL := fmt.Sprintf("http://short.url/%d", hash)
	urlMap[shortURL] = longURL
	return shortURL
}

func redirectToLongURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	longURL, exists := urlMap[shortURL]
	if exists {
		http.Redirect(w, r, longURL, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		longURL := r.FormValue("url")
		shortURL := generateShortenURL(longURL)
		fmt.Fprintf(w, "Short URL: %s", &shortURL)
		http.HandleFunc("/", redirectToLongURL)
		http.ListenAndServe(":8080", nil)
	})
}
