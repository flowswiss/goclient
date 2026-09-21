package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type ConnectionService struct {
	client *core.Client
}

func NewConnectionService(client *core.Client) *ConnectionService {
	return &ConnectionService{client: client}
}

func (c ConnectionService) List(ctx context.Context, cursor core.Cursor) (list common.List[Connection], err error) {
	list.Pagination, err = c.client.List(ctx, getConnectionsPath(), cursor, &list.Items)
	return
}

type ConnectionGetReq struct {
	ID uint `json:"-"`
}

func (c ConnectionService) Get(ctx context.Context, req ConnectionGetReq) (connection Connection, err error) {
	err = c.client.Get(ctx, getSpecificConnectionPath(req.ID), &connection)
	return
}

type ConnectionPeeringCreateReq struct {
	Local  ConnectionEndpoint `json:"local_endpoint"`
	Remote ConnectionEndpoint `json:"remote_endpoint"`
}

func (c ConnectionService) CreatePeering(ctx context.Context, req ConnectionPeeringCreateReq) (
	orderings []common.Ordering,
	err error,
) {
	err = c.client.Create(ctx, getConnectionPeeringPath(), req, &orderings)
	return
}

type ConnectionVPNCreateReq struct {
	Name              string                      `json:"name"`
	PSK               string                      `json:"psk"`
	MTU               int                         `json:"mtu"`
	InitiatorID       int                         `json:"initiator_id"`
	IKEPolicy         ConnectionIKEPolicy         `json:"ike_policy"`
	IPSecPolicy       ConnectionIPSecPolicy       `json:"ipsec_policy"`
	LocalEndpoint     ConnectionEndpoint          `json:"local_endpoint"`
	RemoteEndpoint    ConnectionExternalEndpoint  `json:"remote_endpoint"`
	DeadPeerDetection ConnectionDeadPeerDetection `json:"dead_peer_detection"`
}

func (c ConnectionService) CreateVPN(ctx context.Context, req ConnectionVPNCreateReq) (
	ordering common.Ordering,
	err error,
) {
	err = c.client.Create(ctx, getConnectionVPNPath(), req, &ordering)
	return
}

type ConnectionPerformReq struct {
	ID uint `json:"-"`

	Action string `json:"action"`
}

func (c ConnectionService) Perform(ctx context.Context, req ConnectionPerformReq) (connection Connection, err error) {
	err = c.client.Create(ctx, getConnectionActionPath(req.ID), req, &connection)
	return
}

type ConnectionUpdateReq struct {
	ID uint `json:"-"`

	Name              *string                      `json:"name,omitempty"`
	PSK               *string                      `json:"psk,omitempty"`
	MTU               *int                         `json:"mtu,omitempty"`
	InitiatorID       *int                         `json:"initiator_id,omitempty"`
	LocalEndpoint     *ConnectionEndpoint          `json:"local_endpoint,omitempty"`
	RemoteEndpoint    *ConnectionExternalEndpoint  `json:"remote_endpoint,omitempty"`
	DeadPeerDetection *ConnectionDeadPeerDetection `json:"dead_peer_detection,omitempty"`
}

func (c ConnectionService) Update(ctx context.Context, req ConnectionUpdateReq) (connection Connection, err error) {
	err = c.client.Update(ctx, getSpecificConnectionPath(req.ID), req, &connection)
	return
}

type ConnectionDeleteReq struct {
	ID uint `json:"-"`
}

func (c ConnectionService) Delete(ctx context.Context, req ConnectionDeleteReq) (err error) {
	err = c.client.Delete(ctx, getSpecificConnectionPath(req.ID))
	return
}

const (
	connectionsSegment       = "/v4/compute/connections"
	connectionPeeringSegment = "peering"
	connectionVPNSegment     = "vpn"
	connectionActionSegment  = "action"
)

func getConnectionsPath() string {
	return connectionsSegment
}

func getConnectionPeeringPath() string {
	return core.Join(connectionsSegment, connectionPeeringSegment)
}

func getConnectionVPNPath() string {
	return core.Join(connectionsSegment, connectionVPNSegment)
}

func getSpecificConnectionPath(connectionID uint) string {
	return core.Join(connectionsSegment, connectionID)
}

func getConnectionActionPath(connectionID uint) string {
	return core.Join(connectionsSegment, connectionID, connectionActionSegment)
}
