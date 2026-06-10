package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	healthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello?name=Alice", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	if got := w.Body.String(); got != "Hello, Alice!\n" {
		t.Errorf("got %q, want %q", got, "Hello, Alice!\n")
	}
}

func TestHelloHandlerDefault(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	if got := w.Body.String(); got != "Hello, world!\n" {
		t.Errorf("got %q, want %q", got, "Hello, world!\n")
	}
}
