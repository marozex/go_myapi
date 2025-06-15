package main

import (
	"log"
	"net/http"

	"github.com/marozex/go_myapi/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)

	log.Println("go server start")
	http.ListenAndServe(":1234", nil)
}
