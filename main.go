package main

import (
	"go-messanger/handler"
	"log"
	"net/http"
)

func main() {

	hello := handler.NewHelloHandler()
	bye := handler.NewByeHandler()

	http.Handle("/hello", hello)
	http.Handle("/bye", bye)

	log.Println("API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
