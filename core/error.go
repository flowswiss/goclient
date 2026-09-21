package core

import "net/http"

type APIError struct {
	response  *http.Response
	message   string
	requestID string
}

func (a APIError) Response() *http.Response { return a.response }
func (a APIError) Error() string            { return a.message }
func (a APIError) RequestID() string        { return a.requestID }
