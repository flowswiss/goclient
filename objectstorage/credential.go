package objectstorage

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Credential struct {
	ID        int             `json:"id"`
	Location  common.Location `json:"location"`
	Endpoint  string          `json:"endpoint"`
	AccessKey string          `json:"access_key"`
	SecretKey string          `json:"secret_key"`
}

type CredentialList struct {
	Items      []Credential
	Pagination goclient.Pagination
}

type CredentialService struct {
	client goclient.Client
}

func NewCredentialService(client goclient.Client) CredentialService {
	return CredentialService{
		client: client,
	}
}

func (i CredentialService) List(ctx context.Context, cursor goclient.Cursor) (list CredentialList, err error) {
	list.Pagination, err = i.client.List(ctx, getCredentialSegment(), cursor, &list.Items)
	return
}

const credentialSegment = "/v4/object-storage/credentials"

func getCredentialSegment() string {
	return credentialSegment
}
