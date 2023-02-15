package main

import (
	"log"
	"net/http"

	"github.com/cisagov/con-pca-tasks/controllers"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Print the version and exit if the -version flag is provided
	version()

	mux := chi.NewRouter()
	mux.Mount("/tasks", controllers.TasksRouter())

	port := ":8080"
	log.Printf("listening on port %s, version %s", port, Version)
	log.Println(http.ListenAndServe(port, mux))
}
