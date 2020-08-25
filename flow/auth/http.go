package auth

import "net/http"

func NewClientWithTransport(authenticator Authenticator, transport http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: &Transport{
			Authenticator: authenticator,
			Base:          transport,
		},
	}
}

func NewClient(authenticator Authenticator) *http.Client {
	return NewClientWithTransport(authenticator, http.DefaultTransport)
}

func NewClientWithToken(token string) *http.Client {
	return NewClient(ApplicationTokenAuthenticator{Token: token})
}
