package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/welcome", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WelcomeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned error status code: got %v expected %v",
			status, http.StatusOK)
	}

	expected := `{"message": "Hello from Golang"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v expected %v",
			rr.Body.String(), expected)
	}

}
