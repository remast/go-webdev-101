package main

import (
	"encoding/json"
	"net/http"
)

type Cat struct {
	Name string `json:"name"`
}

func catAPIHandler(w http.ResponseWriter, r *http.Request) {
	c := Cat{Name: "Ginger"}
	json.NewEncoder(w).Encode(c)
}

func main() {
	http.HandleFunc("/api/cat", catAPIHandler)
	http.ListenAndServe(":8080", nil)
}
