package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marozex/go_myapi/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)

	log.Println("server start at 1234")
	http.ListenAndServe(":1234", r)
}
