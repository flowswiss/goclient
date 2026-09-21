package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type KeyPairService struct {
	client *core.Client
}

func NewKeyPairService(client *core.Client) *KeyPairService {
	return &KeyPairService{client: client}
}

func (k KeyPairService) List(ctx context.Context, cursor core.Cursor) (list common.List[KeyPair], err error) {
	list.Pagination, err = k.client.List(ctx, getKeyPairsPath(), cursor, &list.Items)
	return
}

type KeyPairCreateReq struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

func (k KeyPairService) Create(ctx context.Context, req KeyPairCreateReq) (keyPair KeyPair, err error) {
	err = k.client.Create(ctx, getKeyPairsPath(), req, &keyPair)
	return
}

type KeyPairDeleteReq struct {
	ID uint `json:"-"`
}

func (k KeyPairService) Delete(ctx context.Context, req KeyPairDeleteReq) (err error) {
	err = k.client.Delete(ctx, getSpecificKeyPairPath(req.ID))
	return
}

const keyPairsSegment = "/v4/compute/key-pairs"

func getKeyPairsPath() string {
	return keyPairsSegment
}

func getSpecificKeyPairPath(id uint) string {
	return core.Join(keyPairsSegment, id)
}
