package flow

import (
	"fmt"
	"net/http"
)

const (
	ErrorUnsupportedContentType = SystemError("received unsupported content type")
)

type SystemError string

func (e SystemError) Error() string {
	return string(e)
}

type ErrorResponse struct {
	Response  *http.Response
	Message   string
	RequestID string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s %s (request %s) resulted in %s: %s", e.Response.Request.Method, e.Response.Request.URL, e.RequestID, e.Response.Status, e.Message)
}
