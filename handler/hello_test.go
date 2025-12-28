package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	// Act
	helloHandler(rec, req)

	// Assert
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	body := rec.Body.String()
	expected := "{\"message\":\"Hello, world!\"}"
	if body != expected {
		t.Fatalf("expected body %q, got %q", expected, body)
	}
}
