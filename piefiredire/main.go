package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /beef/summary", beefSummary)

	log.Println("starting server on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
