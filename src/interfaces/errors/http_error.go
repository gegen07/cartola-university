package errors

import (
	"encoding/json"
	"fmt"
)

type HTTPError struct {
	Cause error `json:"-"`
	Detail string `json:"detail"`
	Status int `json:"-"`
}

func NewHTTPError(err error, status int, detail string) *HTTPError {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}

	return e.Detail + ":" + e.Cause.Error()
}

func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)

	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}

	return body, nil
}

func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string {
		"Content-Type": "application/json; charset=utf-8",
	}
}