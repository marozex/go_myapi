package handlers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		io.WriteString(w, "hello this is go_myapi")
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}
