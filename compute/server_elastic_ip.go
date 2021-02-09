package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type ElasticIpAttach struct {
	ElasticIpId        int `json:"elastic_ip_id"`
	NetworkInterfaceId int `json:"network_interface_id"`
}

type ServerElasticIpService struct {
	client   goclient.Client
	serverId int
}

func (s ServerElasticIpService) List(ctx context.Context, cursor goclient.Cursor) (list ElasticIpList, err error) {
	list.Pagination, err = s.client.List(ctx, getServerElasticIpsPath(s.serverId), cursor, &list.Items)
	return
}

func (s ServerElasticIpService) Attach(ctx context.Context, body ElasticIpAttach) (elasticIp ElasticIp, err error) {
	err = s.client.Create(ctx, getServerElasticIpsPath(s.serverId), body, &elasticIp)
	return
}

func (s ServerElasticIpService) Detach(ctx context.Context, id int) (err error) {
	err = s.client.Delete(ctx, getSpecificServerElasticIpPath(s.serverId, id))
	return
}

const serverElasticIpsSegment = "elastic-ips"

func getServerElasticIpsPath(serverId int) string {
	return goclient.Join(serversSegment, serverId, serverElasticIpsSegment)
}

func getSpecificServerElasticIpPath(serverId, elasticIpId int) string {
	return goclient.Join(serversSegment, serverId, serverElasticIpsSegment, elasticIpId)
}
