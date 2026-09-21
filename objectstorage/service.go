package objectstorage

import "github.com/flowswiss/goclient/v2/core"

type Service struct {
	Credential *CredentialService
	Instance   *InstanceService
}

func NewService(client *core.Client) *Service {
	return &Service{
		Credential: NewCredentialService(client),
		Instance:   NewInstanceService(client),
	}
}
