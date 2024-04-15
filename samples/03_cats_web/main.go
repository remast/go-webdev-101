package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

type Cat struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       struct {
		URL string `json:"url"`
	} `json:"image"`
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("https://api.thecatapi.com/v1/breeds?api_key=%v&limit=5", os.Getenv("CATS_API_KEY")))
	if err != nil {
		http.Error(w, "Cats API error", http.StatusInternalServerError)
		return
	}

	// Create cat slice
	cat := make([]Cat, 5)

	// Read and parse
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &cat)

	// Render template
	tpl.Execute(w, cat)
}

//go:embed assets
var assets embed.FS

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.FileServer(http.FS(assets)))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/cat", catAPIHandler)

	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
