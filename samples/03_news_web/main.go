package main

import (
	"embed"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Cat struct {
	Name  string `json:"name"`
	Image struct {
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

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// Create cats slice
	cats := make([]Cat, 5)
	json.Unmarshal(body, &cats)

	// Render template
	tpl.Execute(w, cats)
}

//go:embed assets
var assets embed.FS

func main() {

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.FileServer(http.FS(assets)))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
