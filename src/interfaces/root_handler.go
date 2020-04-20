package interfaces

import (
	"github.com/gegen07/cartola-university/interfaces/errors"
	"log"
	"net/http"
)

type RootHandler func(http.ResponseWriter, *http.Request) error

func (fn RootHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)

	if err == nil {
		return
	}

	log.Printf("An error occured: %v", err)

	clientError, ok := err.(errors.ClientError)

	if !ok {
		w.WriteHeader(500)
		return
	}

	body, err := clientError.ResponseBody()

	if err != nil {
		w.WriteHeader(500)
		return
	}

	status, headers := clientError.ResponseHeaders()

	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	w.Write(body)
}