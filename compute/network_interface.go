package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type NetworkInterface struct {
	ID                int             `json:"id"`
	PrivateIP         string          `json:"private_ip"`
	MacAddress        string          `json:"mac_address"`
	Network           Network         `json:"network"`
	AttachedElasticIP ElasticIP       `json:"attached_elastic_ip"`
	SecurityGroups    []SecurityGroup `json:"security_groups"`
	Security          bool            `json:"security"`
}

type NetworkInterfaceList struct {
	Items      []NetworkInterface
	Pagination goclient.Pagination
}

type NetworkInterfaceCreate struct {
	NetworkID int    `json:"network_id"`
	PrivateIP string `json:"private_ip"`
}

type NetworkInterfaceSecurityUpdate struct {
	Security bool `json:"security"`
}

type NetworkInterfaceSecurityGroupUpdate struct {
	SecurityGroupIDs []int `json:"security_group_ids"`
}

type NetworkInterfaceService struct {
	client   goclient.Client
	serverID int
}

func NewNetworkInterfaceService(client goclient.Client, serverID int) NetworkInterfaceService {
	return NetworkInterfaceService{client: client, serverID: serverID}
}

func (n NetworkInterfaceService) List(ctx context.Context, cursor goclient.Cursor) (list NetworkInterfaceList, err error) {
	list.Pagination, err = n.client.List(ctx, getNetworkInterfacesPath(n.serverID), cursor, &list.Items)
	return
}

func (n NetworkInterfaceService) Create(ctx context.Context, body NetworkInterfaceCreate) (networkInterface NetworkInterface, err error) {
	err = n.client.Create(ctx, getNetworkInterfacesPath(n.serverID), body, &networkInterface)
	return
}

func (n NetworkInterfaceService) UpdateSecurity(ctx context.Context, id int, body NetworkInterfaceSecurityUpdate) (networkInterface NetworkInterface, err error) {
	err = n.client.Update(ctx, getNetworkInterfaceSecurityPath(n.serverID, id), body, &networkInterface)
	return
}

func (n NetworkInterfaceService) UpdateSecurityGroups(ctx context.Context, id int, body NetworkInterfaceSecurityGroupUpdate) (networkInterface NetworkInterface, err error) {
	err = n.client.Update(ctx, getNetworkInterfaceSecurityGroupsPath(n.serverID, id), body, &networkInterface)
	return
}

func (n NetworkInterfaceService) Delete(ctx context.Context, id int) (err error) {
	err = n.client.Delete(ctx, getSpecificNetworkInterfacePath(n.serverID, id))
	return
}

const (
	networkInterfacesSegment              = "network-interfaces"
	networkInterfaceSecuritySegment       = "security"
	networkInterfaceSecurityGroupsSegment = "security-groups"
)

func getNetworkInterfacesPath(serverID int) string {
	return goclient.Join(serversSegment, serverID, networkInterfacesSegment)
}

func getSpecificNetworkInterfacePath(serverID, networkInterfaceID int) string {
	return goclient.Join(serversSegment, serverID, networkInterfacesSegment, networkInterfaceID)
}

func getNetworkInterfaceSecurityPath(serverID, networkInterfaceID int) string {
	return goclient.Join(serversSegment, serverID, networkInterfacesSegment, networkInterfaceID, networkInterfaceSecuritySegment)
}

func getNetworkInterfaceSecurityGroupsPath(serverID, networkInterfaceID int) string {
	return goclient.Join(serversSegment, serverID, networkInterfacesSegment, networkInterfaceID, networkInterfaceSecurityGroupsSegment)
}
