package compute

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type ServerService struct {
	client *core.Client
}

func NewServerService(client *core.Client) *ServerService {
	return &ServerService{client: client}
}

func (s ServerService) List(ctx context.Context, cursor core.Cursor) (list common.List[Server], err error) {
	list.Pagination, err = s.client.List(ctx, getServersPath(), cursor, &list.Items)
	return
}

type ServerGetReq struct {
	ID uint `json:"-"`
}

func (s ServerService) Get(ctx context.Context, req ServerGetReq) (server Server, err error) {
	err = s.client.Get(ctx, getSpecificServerPath(req.ID), &server)
	return
}

type ServerCreateReq struct {
	Name             string  `json:"name"`
	LocationID       int     `json:"location_id"`
	ImageID          int     `json:"image_id"`
	ProductID        int     `json:"product_id"`
	AttachExternalIP bool    `json:"attach_external_ip"`
	NetworkID        int     `json:"network_id"`
	PrivateIP        *string `json:"private_ip,omitempty"`
	KeyPairID        *int    `json:"key_pair_id,omitempty"`
	Password         *string `json:"password,omitempty"`
	CloudInit        *string `json:"cloud_init,omitempty"`
}

func (s ServerService) Create(ctx context.Context, req ServerCreateReq) (ordering common.Ordering, err error) {
	err = s.client.Create(ctx, getServersPath(), req, &ordering)
	return
}

type ServerPerformReq struct {
	ID uint `json:"-"`

	Action string `json:"action"`
}

func (s ServerService) Perform(ctx context.Context, req ServerPerformReq) (server Server, err error) {
	err = s.client.Create(ctx, getServerActionPath(req.ID), req, &server)
	return
}

type ServerUpdateReq struct {
	ID uint `json:"-"`

	Name string `json:"name"`
}

func (s ServerService) Update(ctx context.Context, req ServerUpdateReq) (server Server, err error) {
	err = s.client.Update(ctx, getSpecificServerPath(req.ID), req, &server)
	return
}

type ServerUpgradeReq struct {
	ID uint `json:"-"`

	ProductID int `json:"product_id"`
}

func (s ServerService) Upgrade(ctx context.Context, req ServerUpgradeReq) (ordering common.Ordering, err error) {
	err = s.client.Create(ctx, getServerUpgradePath(req.ID), req, &ordering)
	return
}

type ServerDeleteReq struct {
	ID              uint `json:"-"`
	DeleteElasticIP bool `json:"-"`
}

func (s ServerService) Delete(ctx context.Context, req ServerDeleteReq) (err error) {
	query := url.Values{
		"delete_elastic_ip": []string{strconv.FormatBool(req.DeleteElasticIP)},
	}

	path := fmt.Sprint(getSpecificServerPath(req.ID), "?", query.Encode())
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

func getSpecificServerPath(serverID uint) string {
	return core.Join(serversSegment, serverID)
}

func getServerActionPath(serverID uint) string {
	return core.Join(serversSegment, serverID, serverActionSegment)
}

func getServerUpgradePath(serverID uint) string {
	return core.Join(serversSegment, serverID, serverUpgradeSegment)
}
