package main

import (
	"encoding/json"
	"net/http"
)

type CatAPI struct {
	Name string `json:"name"`
}

func catAPIHandler(w http.ResponseWriter, r *http.Request) {
	c := CatAPI{Name: "Ginger"}
	json.NewEncoder(w).Encode(c)
}
