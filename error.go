package main

import "fmt"

// ErrorAPI : Object For Error Handlign within the API
type ErrorAPI struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"httpStatus"`
}

func (e *ErrorAPI) Error() string {
	return fmt.Sprintf("[%d] %s", e.HTTPStatus, e.Message)
}
