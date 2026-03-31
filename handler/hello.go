package handler

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Hello, world!"}`))
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Bye"}`))
}

func NewHelloHandler() http.Handler {
	return http.HandlerFunc(helloHandler)
}

func NewByeHandler() http.Handler {
	return http.HandlerFunc(byeHandler)
}
