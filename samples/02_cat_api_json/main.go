package main

import (
	"encoding/json"
	"net/http"
)

type Cat struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func catAPIHandler(w http.ResponseWriter, r *http.Request) {
	cats := make([]Cat, 1)
	cats[0] = Cat{Name: "Ginger", Image: "https://cdn2.thecatapi.com/images/0XYvRd7oD.jpg"}
	json.NewEncoder(w).Encode(cats)
}

func main() {
	http.HandleFunc("/api/cats", catAPIHandler)
	http.ListenAndServe(":8080", nil)
}
