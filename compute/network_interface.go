package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type NetworkInterfaceService struct {
	client *core.Client
}

func NewNetworkInterfaceService(client *core.Client) *NetworkInterfaceService {
	return &NetworkInterfaceService{client: client}
}

type NetworkInterfaceListReq struct {
	ServerID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (n NetworkInterfaceService) List(ctx context.Context, req NetworkInterfaceListReq) (
	list common.List[NetworkInterface],
	err error,
) {
	list.Pagination, err = n.client.List(ctx, getNetworkInterfacesPath(req.ServerID), req.Cursor, &list.Items)
	return
}

type NetworkInterfaceCreateReq struct {
	ServerID uint `json:"-"`

	NetworkID int    `json:"network_id"`
	PrivateIP string `json:"private_ip"`
}

func (n NetworkInterfaceService) Create(
	ctx context.Context,
	req NetworkInterfaceCreateReq,
) (networkInterface NetworkInterface, err error) {
	err = n.client.Create(ctx, getNetworkInterfacesPath(req.ServerID), req, &networkInterface)
	return
}

type NetworkInterfaceSecurityUpdateReq struct {
	ServerID           uint `json:"-"`
	NetworkInterfaceID uint `json:"-"`

	Security bool `json:"security"`
}

func (n NetworkInterfaceService) UpdateSecurity(
	ctx context.Context,
	req NetworkInterfaceSecurityUpdateReq,
) (networkInterface NetworkInterface, err error) {
	err = n.client.Update(ctx, getNetworkInterfaceSecurityPath(req.ServerID, req.NetworkInterfaceID), req, &networkInterface)
	return
}

type NetworkInterfaceSecurityGroupUpdateReq struct {
	ServerID           uint `json:"-"`
	NetworkInterfaceID uint `json:"-"`

	SecurityGroupIDs []int `json:"security_group_ids"`
}

func (n NetworkInterfaceService) UpdateSecurityGroups(
	ctx context.Context,
	req NetworkInterfaceSecurityGroupUpdateReq,
) (networkInterface NetworkInterface, err error) {
	err = n.client.Update(ctx, getNetworkInterfaceSecurityGroupsPath(req.ServerID, req.NetworkInterfaceID), req, &networkInterface)
	return
}

type NetworkInterfaceDeleteReq struct {
	ServerID           uint `json:"-"`
	NetworkInterfaceID uint `json:"-"`
}

func (n NetworkInterfaceService) Delete(ctx context.Context, req NetworkInterfaceDeleteReq) (err error) {
	err = n.client.Delete(ctx, getSpecificNetworkInterfacePath(req.ServerID, req.NetworkInterfaceID))
	return
}

const (
	networkInterfacesSegment              = "network-interfaces"
	networkInterfaceSecuritySegment       = "security"
	networkInterfaceSecurityGroupsSegment = "security-groups"
)

func getNetworkInterfacesPath(serverID uint) string {
	return core.Join(serversSegment, serverID, networkInterfacesSegment)
}

func getSpecificNetworkInterfacePath(serverID, networkInterfaceID uint) string {
	return core.Join(serversSegment, serverID, networkInterfacesSegment, networkInterfaceID)
}

func getNetworkInterfaceSecurityPath(serverID, networkInterfaceID uint) string {
	return core.Join(serversSegment, serverID, networkInterfacesSegment, networkInterfaceID, networkInterfaceSecuritySegment)
}

func getNetworkInterfaceSecurityGroupsPath(serverID, networkInterfaceID uint) string {
	return core.Join(serversSegment, serverID, networkInterfacesSegment, networkInterfaceID, networkInterfaceSecurityGroupsSegment)
}
