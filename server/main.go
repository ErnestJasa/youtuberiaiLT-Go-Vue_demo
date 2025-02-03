package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/youtubers", GetChannels)
	mux.HandleFunc("/api/categories", GetCategories)

	handlerWithCORS := EnableCORS(mux)
	log.Println("server is running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))
}
