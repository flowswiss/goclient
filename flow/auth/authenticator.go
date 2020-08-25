package auth

type Authenticator interface {
	GetToken() string
}

type ApplicationTokenAuthenticator struct {
	Token string
}

func (a ApplicationTokenAuthenticator) GetToken() string {
	return a.Token
}
