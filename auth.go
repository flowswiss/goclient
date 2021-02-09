package goclient

import (
	"fmt"
	"net/http"
)

type AuthTransport struct {
	Token string
	Base  http.RoundTripper
}

func (t AuthTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.Token))

	return t.base().RoundTrip(request)
}

func (t AuthTransport) base() http.RoundTripper {
	if t.Base == nil {
		return http.DefaultTransport
	}

	return t.Base
}
