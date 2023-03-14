package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cisagov/con-pca-tasks/controllers"
	db "github.com/cisagov/con-pca-tasks/database"
	"github.com/cisagov/con-pca-tasks/notifications"
	"github.com/cisagov/con-pca-tasks/services/aws"
	"github.com/go-chi/chi/v5"
)

var apiKey string

func init() {
	// Load the API key from the environment
	apiKey = os.Getenv("API_ACCESS_KEY")
	notifications.ApiUrl = os.Getenv("API_URL")

	// Connect to the database
	db.InitDB()

	// Initialize SES email
	aws.SESEmailClient()
}

// Auth is a middleware that verifies the API key
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if apiKey != r.Header.Get("api_key") {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return

		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Print the version and exit if the -version flag is provided
	version()

	r := chi.NewRouter()

	// Protected routes
	r.Group(func(r chi.Router) {
		// Verify API Key
		r.Use(Auth)

		// Mount the tasks route to the router
		r.Mount("/tasks", controllers.TasksRouter())
	})

	port := ":8080"
	log.Printf("listening on port %s, version %s", port, Version)
	log.Println(http.ListenAndServe(port, r))
}
