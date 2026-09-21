package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type ElasticIPService struct {
	client *core.Client
}

func NewElasticIPService(client *core.Client) *ElasticIPService {
	return &ElasticIPService{client: client}
}

func (e ElasticIPService) List(ctx context.Context, cursor core.Cursor) (list common.List[ElasticIP], err error) {
	list.Pagination, err = e.client.List(ctx, getElasticIPsPath(), cursor, &list.Items)
	return
}

type ElasticIPCreateReq struct {
	LocationID int `json:"location_id,omitempty"`
}

func (e ElasticIPService) Create(ctx context.Context, req ElasticIPCreateReq) (elasticIP ElasticIP, err error) {
	err = e.client.Create(ctx, getElasticIPsPath(), req, &elasticIP)
	return
}

type ElasticIPDeleteReq struct {
	ID uint `json:"-"`
}

func (e ElasticIPService) Delete(ctx context.Context, req ElasticIPDeleteReq) (err error) {
	err = e.client.Delete(ctx, getSpecificElasticIPPath(req.ID))
	return
}

type ElasticIPAttachmentService struct {
	client *core.Client
}

func NewElasticIPAttachmentService(client *core.Client) *ElasticIPAttachmentService {
	return &ElasticIPAttachmentService{client: client}
}

type ElasticIPAttachmentListReq struct {
	DeviceID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (a ElasticIPAttachmentService) List(ctx context.Context, req ElasticIPAttachmentListReq) (
	list common.List[ElasticIP],
	err error,
) {
	list.Pagination, err = a.client.List(ctx, getAttachedElasticIPsPath(req.DeviceID), req.Cursor, &list.Items)
	return
}

type ElasticIPAttachmentCreateReq struct {
	DeviceID uint `json:"-"`

	ElasticIPID        int `json:"elastic_ip_id"`
	NetworkInterfaceID int `json:"network_interface_id"`
}

func (a ElasticIPAttachmentService) Create(ctx context.Context, req ElasticIPAttachmentCreateReq) (
	elasticIP ElasticIP,
	err error,
) {
	err = a.client.Create(ctx, getAttachedElasticIPsPath(req.DeviceID), req, &elasticIP)
	return
}

type ElasticIPAttachmentDeleteReq struct {
	DeviceID    uint `json:"-"`
	ElasticIPID uint `json:"-"`
}

func (a ElasticIPAttachmentService) Delete(ctx context.Context, req ElasticIPAttachmentDeleteReq) (err error) {
	err = a.client.Delete(ctx, getSpecificAttachedElasticIPPath(req.DeviceID, req.ElasticIPID))
	return
}

const (
	elasticIPsSegment         = "/v4/macbaremetal/elastic-ips"
	attachedElasticIPsSegment = "elastic-ips"
)

func getElasticIPsPath() string {
	return elasticIPsSegment
}

func getSpecificElasticIPPath(id uint) string {
	return core.Join(elasticIPsSegment, id)
}

func getAttachedElasticIPsPath(deviceID uint) string {
	return core.Join(getSpecificDevicePath(deviceID), attachedElasticIPsSegment)
}

func getSpecificAttachedElasticIPPath(deviceID, elasticIPID uint) string {
	return core.Join(getAttachedElasticIPsPath(deviceID), elasticIPID)
}
