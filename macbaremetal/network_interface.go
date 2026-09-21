package macbaremetal

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
	DeviceID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (n NetworkInterfaceService) List(ctx context.Context, req NetworkInterfaceListReq) (
	list common.List[NetworkInterface],
	err error,
) {
	list.Pagination, err = n.client.List(ctx, getNetworkInterfacesPath(req.DeviceID), req.Cursor, &list.Items)
	return
}

type NetworkInterfaceSecurityGroupUpdateReq struct {
	DeviceID           uint `json:"-"`
	NetworkInterfaceId uint `json:"-"`

	SecurityGroupID int `json:"security_group_id"`
}

func (n NetworkInterfaceService) UpdateSecurityGroup(
	ctx context.Context,
	req NetworkInterfaceSecurityGroupUpdateReq,
) (networkInterface NetworkInterface, err error) {
	err = n.client.Update(ctx, getSpecificNetworkInterfacePath(req.DeviceID, req.NetworkInterfaceId), req, &networkInterface)
	return
}

const networkInterfacesSegment = "network-interfaces"

func getNetworkInterfacesPath(deviceID uint) string {
	return core.Join(getSpecificDevicePath(deviceID), networkInterfacesSegment)
}

func getSpecificNetworkInterfacePath(deviceID, networkInterfaceID uint) string {
	return core.Join(getNetworkInterfacesPath(deviceID), networkInterfaceID)
}
