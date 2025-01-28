package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/youtubers", GetChannels)
	http.HandleFunc("/api/categories", GetCategories)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
