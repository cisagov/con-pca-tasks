package main

import (
	"log"
	"net/http"

	"github.com/cisagov/con-pca-tasks/controllers"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", controllers.HealthCheckHandler)

	port := ":8080"
	log.Printf("listening on port %s", port)
	log.Println(http.ListenAndServe(port, mux))
}
