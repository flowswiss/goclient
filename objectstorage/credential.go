package objectstorage

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type Credential struct {
	ID        int             `json:"id"`
	Location  common.Location `json:"location"`
	Endpoint  string          `json:"endpoint"`
	AccessKey string          `json:"access_key"`
	SecretKey string          `json:"secret_key"`
}

type CredentialService struct {
	client *core.Client
}

func NewCredentialService(client *core.Client) *CredentialService {
	return &CredentialService{
		client: client,
	}
}

func (i CredentialService) List(ctx context.Context, cursor core.Cursor) (list common.List[Credential], err error) {
	list.Pagination, err = i.client.List(ctx, getCredentialSegment(), cursor, &list.Items)
	return
}

const credentialSegment = "/v4/object-storage/credentials"

func getCredentialSegment() string {
	return credentialSegment
}
