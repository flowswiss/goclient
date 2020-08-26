package auth

import (
	"fmt"
	"net/http"
)

type Transport struct {
	Authenticator Authenticator
	Base          http.RoundTripper
}

func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	token := t.Authenticator.GetToken()
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return t.Base.RoundTrip(request)
}
