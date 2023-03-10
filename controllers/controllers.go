package controllers

import (
	"fmt"
	"net/http"

	"github.com/cisagov/con-pca-tasks/notifications"
	"github.com/go-chi/chi/v5"
)

// healthCheckHandler indicates that the server is up and running.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	println("Health check reached.")
	fmt.Fprintf(w, "Up and running!")
}

// emailReportHandler manages emailing notification reports.
func emailReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cycleId := chi.URLParam(r, "cycle_id")
	reportType := chi.URLParam(r, "report_type")

	notifications.Manager(cycleId, reportType)
	fmt.Fprintf(w, "%s report email sent! Cycle id: %s", reportType, cycleId)
}

// pdfReportHandler manages PDF reports.
func pdfReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cycleId := chi.URLParam(r, "cycle_id")
	reportType := chi.URLParam(r, "report_type")

	fmt.Fprintf(w, "%s report pdf complete! Cycle id: %s", reportType, cycleId)
}

func TasksRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", healthCheckHandler)
	r.Get("/{cycle_id}/reports/{report_type}/email", emailReportHandler)
	r.Get("/{cycle_id}/reports/{report_type}/pdf", pdfReportHandler)

	return r
}
