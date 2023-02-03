package main

import (
	"log"
	"net/http"

	"github.com/cisagov/con-pca-tasks/controllers"
	"github.com/go-chi/chi"
)

func main() {
	// Print the version and exit if the -version flag is provided
	version()

	mux := chi.NewRouter()
	mux.Get("/", controllers.HealthCheckHandler)

	port := ":8080"
	log.Printf("listening on port %s", port)
	log.Println(http.ListenAndServe(port, mux))
}
