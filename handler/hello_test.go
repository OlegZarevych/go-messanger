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

func TestByeHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/bye", nil)
	rec := httptest.NewRecorder()

	// Act
	byeHandler(rec, req)

	// Assert
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", ct)
	}

	body := rec.Body.String()
	expected := "{\"message\":\"Bye\"}"
	if body != expected {
		t.Fatalf("expected body %q, got %q", expected, body)
	}
}
