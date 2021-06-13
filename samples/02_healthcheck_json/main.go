package main

import (
	"encoding/json"
	"net/http"
)

type CheckResult struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	c := CheckResult{Status: "Up and Running"}
	json.NewEncoder(w).Encode(c)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}
