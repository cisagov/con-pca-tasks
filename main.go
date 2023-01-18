package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HealthCheckHandler)

	port := ":8080"
	log.Printf("listening on port %s", port)
	log.Println(http.ListenAndServe(port, mux))
}
