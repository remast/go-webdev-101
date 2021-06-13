package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v requested URL %v", r.Host, r.URL)
		next.ServeHTTP(w, r)
	})
}
