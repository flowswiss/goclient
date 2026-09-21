package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type DeviceService struct {
	client *core.Client
}

func NewDeviceService(client *core.Client) *DeviceService {
	return &DeviceService{client: client}
}

func (d DeviceService) List(ctx context.Context, cursor core.Cursor) (list common.List[Device], err error) {
	list.Pagination, err = d.client.List(ctx, getDevicesPath(), cursor, &list.Items)
	return
}

type DeviceGetReq struct {
	ID uint `json:"-"`
}

func (d DeviceService) Get(ctx context.Context, req DeviceGetReq) (device Device, err error) {
	err = d.client.Get(ctx, getSpecificDevicePath(req.ID), &device)
	return
}

func (d DeviceService) GetVNC(ctx context.Context, req DeviceGetReq) (vnc DeviceVNCConnection, err error) {
	err = d.client.Get(ctx, getDeviceVNCPath(req.ID), &vnc)
	return
}

type DeviceCreateReq struct {
	Name            string `json:"name"`
	LocationID      int    `json:"location_id"`
	ProductID       int    `json:"product_id"`
	NetworkID       int    `json:"network_id"`
	AttachElasticIP bool   `json:"attach_elastic_ip"`
	Password        string `json:"password"`
}

func (d DeviceService) Create(ctx context.Context, req DeviceCreateReq) (order common.Ordering, err error) {
	err = d.client.Create(ctx, getDevicesPath(), req, &order)
	return
}

type DeviceUpdateReq struct {
	ID uint `json:"-"`

	Name string `json:"name"`
}

func (d DeviceService) Update(ctx context.Context, req DeviceUpdateReq) (device Device, err error) {
	err = d.client.Update(ctx, getSpecificDevicePath(req.ID), req, &device)
	return
}

type DeviceDeleteReq struct {
	ID uint `json:"-"`
}

func (d DeviceService) Delete(ctx context.Context, req DeviceDeleteReq) (err error) {
	err = d.client.Delete(ctx, getSpecificDevicePath(req.ID))
	return
}

type DevicePerformReq struct {
	DeviceID uint `json:"-"`

	Action string `json:"action"`
}

func (d DeviceService) Perform(ctx context.Context, req DevicePerformReq) (device Device, err error) {
	err = d.client.Create(ctx, getDeviceActionPath(req.DeviceID), req, &device)
	return
}

type DeviceWorkflowListReq struct {
	DeviceID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (d DeviceService) WorkflowList(ctx context.Context, req DeviceWorkflowListReq) (
	list common.List[DeviceWorkflow],
	err error,
) {
	list.Pagination, err = d.client.List(ctx, getDeviceWorkflowPath(req.DeviceID), req.Cursor, &list.Items)
	return
}

type DeviceWorkflowRunReq struct {
	DeviceID uint `json:"-"`

	Workflow string `json:"workflow"`
}

func (d DeviceService) WorkflowRun(ctx context.Context, req DeviceWorkflowRunReq) (device Device, err error) {
	err = d.client.Create(ctx, getDeviceWorkflowPath(req.DeviceID), req, &device)
	return
}

const (
	devicesSegment        = "/v4/macbaremetal/devices"
	deviceVNCSegment      = "vnc"
	deviceActionSegment   = "actions"
	deviceWorkflowSegment = "workflows"
)

func getDevicesPath() string {
	return devicesSegment
}

func getSpecificDevicePath(id uint) string {
	return core.Join(devicesSegment, id)
}

func getDeviceVNCPath(id uint) string {
	return core.Join(getSpecificDevicePath(id), deviceVNCSegment)
}

func getDeviceActionPath(id uint) string {
	return core.Join(devicesSegment, id, deviceActionSegment)
}

func getDeviceWorkflowPath(id uint) string {
	return core.Join(devicesSegment, id, deviceWorkflowSegment)
}
