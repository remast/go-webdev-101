package main

import (
	"embed"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Cat struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       struct {
		URL string `json:"url"`
	}
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds?api_key=c737a14e-40c2-4dca-995d-a87849c4547d&limit=5")
	if err != nil {
		http.Error(w, "Cats API error", http.StatusInternalServerError)
		return
	}

	// Create cat slice
	cat := make([]Cat, 5)

	// Read and parse
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &cat)

	// Render template
	tpl.Execute(w, cat)
}

//go:embed assets
var assets embed.FS

func main() {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.FileServer(http.FS(assets)))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/cat", catAPIHandler)

	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
