package handlers

import (
	"io"
	"net/http"
)

// WelcomeHandler say Hello from golang
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"message": "Hello from Golang"}`)
}
