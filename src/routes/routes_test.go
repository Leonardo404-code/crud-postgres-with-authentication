package routes

import (
	"crud-postgres/src/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	config.LoadDotEnvTests()

	r := Routes()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("error in setup routes want: %v, got: %v", http.StatusOK, w.Code)
	}
}
