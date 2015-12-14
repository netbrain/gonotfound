package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/_ping" {
			w.Write([]byte("pong!"))
			return
		}
		http.NotFound(w, r)
	})))
}
