package main

import (
	"fmt"
	"net/http"
)

func catAPIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Meow!")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/cat", catAPIHandler)
	http.ListenAndServe(":8090", nil)
}
