package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/channels", GetChannels)

	log.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
