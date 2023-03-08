package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	healthCheckHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "Up and running!" {
		t.Errorf("expected 'Up and running!', but got %v", string(data))
	}
}

// func TestEmailReportHandler(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
// 	w := httptest.NewRecorder()

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("cycle_id", "test-cycle-id")
// 	rctx.URLParams.Add("report_type", "cycle")

// 	emailReportHandler(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 	}
// 	if string(data) != "cycle report email sent! Cycle id: test-cycle-id" {
// 		t.Errorf("expected 'cycle report email sent! Cycle id: test-cycle-id', but got %v", string(data))
// 	}
// }
