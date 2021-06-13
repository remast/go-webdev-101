package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Up and running!")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8090", nil)
}
