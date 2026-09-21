package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type ElasticIPAttachmentService struct {
	client *core.Client
}

func NewElasticIPAttachmentService(client *core.Client) *ElasticIPAttachmentService {
	return &ElasticIPAttachmentService{
		client: client,
	}
}

type ElasticIPAttachmentListReq struct {
	ServerID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (s ElasticIPAttachmentService) List(ctx context.Context, req ElasticIPAttachmentListReq) (
	list common.List[ElasticIP],
	err error,
) {
	list.Pagination, err = s.client.List(ctx, getServerElasticIPsPath(req.ServerID), req.Cursor, &list.Items)
	return
}

type ElasticIPAttachmentCreateReq struct {
	ServerID uint `json:"-"`

	ElasticIPID        int `json:"elastic_ip_id"`
	NetworkInterfaceID int `json:"network_interface_id"`
}

func (s ElasticIPAttachmentService) Create(ctx context.Context, req ElasticIPAttachmentCreateReq) (
	elasticIP ElasticIP,
	err error,
) {
	err = s.client.Create(ctx, getServerElasticIPsPath(req.ServerID), req, &elasticIP)
	return
}

type ElasticIPAttachmentDeleteReq struct {
	ServerID    uint `json:"-"`
	ElasticIPID uint `json:"-"`
}

func (s ElasticIPAttachmentService) Delete(ctx context.Context, req ElasticIPAttachmentDeleteReq) (err error) {
	err = s.client.Delete(ctx, getSpecificServerElasticIPPath(req.ServerID, req.ElasticIPID))
	return
}

const serverElasticIPsSegment = "elastic-ips"

func getServerElasticIPsPath(serverID uint) string {
	return core.Join(serversSegment, serverID, serverElasticIPsSegment)
}

func getSpecificServerElasticIPPath(serverID, elasticIPID uint) string {
	return core.Join(serversSegment, serverID, serverElasticIPsSegment, elasticIPID)
}
