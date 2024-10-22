package main

import (
	"log"
	"net/http"
)

const (
	port    = ":8080"
	beefUrl = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /beef/summary", beefSummary)

	log.Println("starting server on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
