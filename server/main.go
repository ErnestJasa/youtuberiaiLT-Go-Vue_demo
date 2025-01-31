package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/youtubers", GetChannels)
	mux.HandleFunc("/api/categories", GetCategories)

	handlerWithCORS := EnableCORS(mux)
	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
