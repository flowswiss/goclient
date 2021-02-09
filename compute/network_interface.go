package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type NetworkInterface struct {
	Id                int             `json:"id"`
	PrivateIp         string          `json:"private_ip"`
	MacAddress        string          `json:"mac_address"`
	Network           Network         `json:"network"`
	AttachedElasticIp ElasticIp       `json:"attached_elastic_ip"`
	SecurityGroups    []SecurityGroup `json:"security_groups"`
	Security          bool            `json:"security"`
}

type NetworkInterfaceList struct {
	Items      []NetworkInterface
	Pagination goclient.Pagination
}

type NetworkInterfaceCreate struct {
	NetworkId int    `json:"network_id"`
	PrivateIp string `json:"private_ip"`
}

type NetworkInterfaceSecurityUpdate struct {
	Security bool `json:"security"`
}

type NetworkInterfaceSecurityGroupUpdate struct {
	SecurityGroupIds []int `json:"security_group_ids"`
}

type NetworkInterfaceService struct {
	client   goclient.Client
	serverId int
}

func NewNetworkInterfaceService(client goclient.Client, serverId int) NetworkInterfaceService {
	return NetworkInterfaceService{client: client, serverId: serverId}
}

func (n NetworkInterfaceService) List(ctx context.Context, cursor goclient.Cursor) (list NetworkInterfaceList, err error) {
	list.Pagination, err = n.client.List(ctx, getNetworkInterfacesPath(n.serverId), cursor, &list.Items)
	return
}

func (n NetworkInterfaceService) Create(ctx context.Context, body NetworkInterfaceCreate) (networkInterface NetworkInterface, err error) {
	err = n.client.Create(ctx, getNetworkInterfacesPath(n.serverId), body, &networkInterface)
	return
}

func (n NetworkInterfaceService) UpdateSecurity(ctx context.Context, id int, body NetworkInterfaceSecurityUpdate) (networkInterface NetworkInterface, err error) {
	err = n.client.Update(ctx, getNetworkInterfaceSecurityPath(n.serverId, id), body, &networkInterface)
	return
}

func (n NetworkInterfaceService) UpdateSecurityGroups(ctx context.Context, id int, body NetworkInterfaceSecurityGroupUpdate) (networkInterface NetworkInterface, err error) {
	err = n.client.Update(ctx, getNetworkInterfaceSecurityGroupsPath(n.serverId, id), body, &networkInterface)
	return
}

func (n NetworkInterfaceService) Delete(ctx context.Context, id int) (err error) {
	err = n.client.Delete(ctx, getSpecificNetworkInterfacePath(n.serverId, id))
	return
}

const (
	networkInterfacesSegment              = "network-interface"
	networkInterfaceSecuritySegment       = "security"
	networkInterfaceSecurityGroupsSegment = "security-groups"
)

func getNetworkInterfacesPath(serverId int) string {
	return goclient.Join(serversSegment, serverId, networkInterfacesSegment)
}

func getSpecificNetworkInterfacePath(serverId, networkInterfaceId int) string {
	return goclient.Join(serversSegment, serverId, networkInterfacesSegment, networkInterfaceId)
}

func getNetworkInterfaceSecurityPath(serverId, networkInterfaceId int) string {
	return goclient.Join(serversSegment, serverId, networkInterfacesSegment, networkInterfaceId, networkInterfaceSecuritySegment)
}

func getNetworkInterfaceSecurityGroupsPath(serverId, networkInterfaceId int) string {
	return goclient.Join(serversSegment, serverId, networkInterfacesSegment, networkInterfaceId, networkInterfaceSecurityGroupsSegment)
}
