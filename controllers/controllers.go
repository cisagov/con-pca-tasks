package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cisagov/con-pca-tasks/aws"
	"github.com/go-chi/chi/v5"
)

// healthCheckHandler indicates that the server is up and running.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	println("Health check reached.")
	fmt.Fprintf(w, "Up and running!")
}

// notificationEmailsHandler manages notificaiton emails.
func notificationEmailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var e aws.SESEmail
	decoder.Decode(&e)

	email := aws.NewSESEmail()
	email.BuildMessage(e.To, e.Cc, e.Bcc, e.Subject, e.Body)
	email.Send()

	fmt.Fprintf(w, "Notification email sent!")
}

func TasksRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", healthCheckHandler)
	r.Post("/notifications/send", notificationEmailsHandler)

	return r
}
