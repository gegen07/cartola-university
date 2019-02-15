package main

import (
	"log"

	"net/http"
	"github.com/gegen07/cartola-coltec/app/cartola-coltec-service/handlers"
	"github.com/gorilla/mux"
)

func main () {
	router := mux.NewRouter()

	router.HandleFunc("/welcome", handlers.WelcomeHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}