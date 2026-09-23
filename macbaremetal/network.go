package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type NetworkService struct {
	client *core.Client
}

func NewNetworkService(client *core.Client) *NetworkService {
	return &NetworkService{client: client}
}

func (n NetworkService) List(ctx context.Context, cursor core.Cursor) (list common.List[Network], err error) {
	list.Pagination, err = n.client.List(ctx, getNetworksPath(), cursor, &list.Items)
	return
}

type NetworkGetReq struct {
	ID uint `json:"-"`
}

func (n NetworkService) Get(ctx context.Context, req NetworkGetReq) (network Network, err error) {
	err = n.client.Get(ctx, getSpecificNetworkPath(req.ID), &network)
	return
}

type NetworkCreateReq struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	LocationID  int     `json:"location_id"`
}

func (n NetworkService) Create(ctx context.Context, req NetworkCreateReq) (network Network, err error) {
	err = n.client.Create(ctx, getNetworksPath(), req, &network)
	return
}

type NetworkUpdateReq struct {
	ID uint `json:"-"`

	Name              *string  `json:"name,omitempty"`
	Description       *string  `json:"description,omitempty"`
	DomainName        *string  `json:"domain_name,omitempty"`
	DomainNameServers []string `json:"domain_name_servers,omitempty"`
}

func (n NetworkService) Update(ctx context.Context, req NetworkUpdateReq) (network Network, err error) {
	err = n.client.Update(ctx, getSpecificNetworkPath(req.ID), req, &network)
	return
}

type NetworkDeleteReq struct {
	ID uint `json:"-"`
}

func (n NetworkService) Delete(ctx context.Context, req NetworkDeleteReq) (err error) {
	err = n.client.Delete(ctx, getSpecificNetworkPath(req.ID))
	return
}

const networksSegment = "/v4/macbaremetal/networks"

func getNetworksPath() string {
	return networksSegment
}

func getSpecificNetworkPath(id uint) string {
	return core.Join(networksSegment, id)
}
