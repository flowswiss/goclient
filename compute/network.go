package compute

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Network struct {
	Id                  int             `json:"id"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	Cidr                string          `json:"cidr"`
	Location            common.Location `json:"location"`
	DomainNameServers   []string        `json:"domain_name_servers"`
	AllocationPoolStart string          `json:"allocation_pool_start"`
	AllocationPoolEnd   string          `json:"allocation_pool_end"`
	GatewayIp           string          `json:"gateway_ip"`
	UsedIps             int             `json:"used_ips"`
	TotalIps            int             `json:"total_ips"`
}

type NetworkList struct {
	Items      []Network
	Pagination goclient.Pagination
}

type NetworkCreate struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	LocationId          int      `json:"location_id"`
	DomainNameServer    []string `json:"domain_name_server"`
	Cidr                string   `json:"cidr"`
	AllocationPoolStart string   `json:"allocation_pool_start"`
	AllocationPoolEnd   string   `json:"allocation_pool_end"`
	GatewayIp           string   `json:"gateway_ip"`
}

type NetworkUpdate struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	DomainNameServer    []string `json:"domain_name_server"`
	AllocationPoolStart string   `json:"allocation_pool_start"`
	AllocationPoolEnd   string   `json:"allocation_pool_end"`
	GatewayIp           string   `json:"gateway_ip"`
}

type NetworkService struct {
	client goclient.Client
}

func NewNetworkService(client goclient.Client) NetworkService {
	return NetworkService{client: client}
}

func (n NetworkService) List(ctx context.Context, cursor goclient.Cursor) (list NetworkList, err error) {
	list.Pagination, err = n.client.List(ctx, getNetworksPath(), cursor, &list.Items)
	return
}

func (n NetworkService) Get(ctx context.Context, id int) (network Network, err error) {
	err = n.client.Get(ctx, getSpecificNetworkPath(id), &network)
	return
}

func (n NetworkService) Create(ctx context.Context, body NetworkCreate) (network Network, err error) {
	err = n.client.Create(ctx, getNetworksPath(), body, &network)
	return
}

func (n NetworkService) Update(ctx context.Context, id int, body NetworkUpdate) (network Network, err error) {
	err = n.client.Update(ctx, getSpecificNetworkPath(id), body, &network)
	return
}

func (n NetworkService) Delete(ctx context.Context, id int) (err error) {
	err = n.client.Delete(ctx, getSpecificNetworkPath(id))
	return
}

const networksSegment = "/v4/compute/networks"

func getNetworksPath() string {
	return networksSegment
}

func getSpecificNetworkPath(networkId int) string {
	return goclient.Join(networksSegment, networkId)
}
