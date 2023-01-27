package controllers

import (
	"fmt"
	"net/http"
)

// HealthCheckHandler indicates that the server is up and running.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Up and running!")
}
