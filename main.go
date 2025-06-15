package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello this is go_myapi")
	}

	http.HandleFunc("/hello", helloHandler)

	log.Println("go server start")
	http.ListenAndServe(":1234", nil)
}
