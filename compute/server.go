package compute

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Server struct {
	ID       int                       `json:"id"`
	Name     string                    `json:"name"`
	Status   ServerStatus              `json:"status"`
	Image    Image                     `json:"image"`
	Product  common.Product            `json:"product"`
	Location common.Location           `json:"location"`
	Networks []ServerNetworkAttachment `json:"networks"`
	KeyPair  KeyPair                   `json:"key_pair"`
}

type ServerList struct {
	Items      []Server
	Pagination goclient.Pagination
}

type ServerNetworkAttachment struct {
	Network
	Interfaces []AttachedNetworkInterface `json:"network_interfaces"`
}

type AttachedNetworkInterface struct {
	ID        int    `json:"id"`
	PrivateIP string `json:"private_ip"`
	PublicIP  string `json:"public_ip"`
}

type ServerCreate struct {
	Name             string `json:"name"`
	LocationID       int    `json:"location_id"`
	ImageID          int    `json:"image_id"`
	ProductID        int    `json:"product_id"`
	AttachExternalIP bool   `json:"attach_external_ip"`
	NetworkID        int    `json:"network_id"`
	PrivateIP        string `json:"private_ip,omitempty"`
	KeyPairID        int    `json:"key_pair_id,omitempty"`
	Password         string `json:"password,omitempty"`
	CloudInit        string `json:"cloud_init,omitempty"`
}

type ServerUpdate struct {
	Name string `json:"name"`
}

type ServerPerform struct {
	Action string `json:"action"`
}

type ServerUpgrade struct {
	ProductID int `json:"product_id"`
}

type ServerService struct {
	client goclient.Client
}

func NewServerService(client goclient.Client) ServerService {
	return ServerService{client: client}
}

func (s ServerService) NetworkInterfaces(serverID int) NetworkInterfaceService {
	return NewNetworkInterfaceService(s.client, serverID)
}

func (s ServerService) List(ctx context.Context, cursor goclient.Cursor) (list ServerList, err error) {
	list.Pagination, err = s.client.List(ctx, getServersPath(), cursor, &list.Items)
	return
}

func (s ServerService) Get(ctx context.Context, id int) (server Server, err error) {
	err = s.client.Get(ctx, getSpecificServerPath(id), &server)
	return
}

func (s ServerService) Create(ctx context.Context, body ServerCreate) (ordering common.Ordering, err error) {
	err = s.client.Create(ctx, getServersPath(), body, &ordering)
	return
}

func (s ServerService) Perform(ctx context.Context, id int, body ServerPerform) (server Server, err error) {
	err = s.client.Create(ctx, getServerActionPath(id), body, &server)
	return
}

func (s ServerService) Update(ctx context.Context, id int, body ServerUpdate) (server Server, err error) {
	err = s.client.Update(ctx, getSpecificServerPath(id), body, &server)
	return
}

func (s ServerService) Upgrade(ctx context.Context, id int, body ServerUpgrade) (ordering common.Ordering, err error) {
	err = s.client.Create(ctx, getServerUpgradePath(id), body, &ordering)
	return
}

func (s ServerService) Delete(ctx context.Context, id int, deleteElasticIP bool) (err error) {
	query := url.Values{
		"delete_elastic_ip": []string{strconv.FormatBool(deleteElasticIP)},
	}

	path := fmt.Sprint(getSpecificServerPath(id), "?", query.Encode())
	err = s.client.Delete(ctx, path)
	return
}

const (
	serversSegment       = "/v4/compute/instances"
	serverActionSegment  = "action"
	serverUpgradeSegment = "upgrade"
)

func getServersPath() string {
	return serversSegment
}

func getSpecificServerPath(serverID int) string {
	return goclient.Join(serversSegment, serverID)
}

func getServerActionPath(serverID int) string {
	return goclient.Join(serversSegment, serverID, serverActionSegment)
}

func getServerUpgradePath(serverID int) string {
	return goclient.Join(serversSegment, serverID, serverUpgradeSegment)
}
