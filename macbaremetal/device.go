package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Device struct {
	ID                int                        `json:"id"`
	Name              string                     `json:"name"`
	Location          common.Location            `json:"location"`
	Product           common.Product             `json:"product"`
	Status            DeviceStatus               `json:"status"`
	OperatingSystem   DeviceOperatingSystem      `json:"operating_system"`
	Network           Network                    `json:"network"`
	Hostname          string                     `json:"hostname"`
	NetworkInterfaces []AttachedNetworkInterface `json:"network_interfaces"`
	Price             float64                    `json:"price"`
	MetalControl      string                     `json:"metal_control"`
	MetalControlTools string                     `json:"metal_control_tools"`
}

type DeviceOperatingSystem struct {
	OS      string `json:"os"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AttachedNetworkInterface struct {
	ID        int    `json:"id"`
	PrivateIP string `json:"private_ip"`
	PublicIP  string `json:"public_ip"`
}

type DeviceList struct {
	Items      []Device
	Pagination goclient.Pagination
}

type DeviceCreate struct {
	Name            string `json:"name"`
	LocationID      int    `json:"location_id"`
	ProductID       int    `json:"product_id"`
	NetworkID       int    `json:"network_id"`
	AttachElasticIP bool   `json:"attach_elastic_ip"`
	Password        string `json:"password"`
}

type DeviceUpdate struct {
	Name string `json:"name,omitempty"`
}

type DeviceService struct {
	client goclient.Client
}

func NewDeviceService(client goclient.Client) DeviceService {
	return DeviceService{client: client}
}

func (d DeviceService) List(ctx context.Context, cursor goclient.Cursor) (list DeviceList, err error) {
	list.Pagination, err = d.client.List(ctx, getDevicesPath(), cursor, &list.Items)
	return
}

func (d DeviceService) Get(ctx context.Context, id int) (device Device, err error) {
	err = d.client.Get(ctx, getSpecificDevicePath(id), &device)
	return
}

func (d DeviceService) Create(ctx context.Context, body DeviceCreate) (order common.Ordering, err error) {
	err = d.client.Create(ctx, getDevicesPath(), body, &order)
	return
}

func (d DeviceService) Update(ctx context.Context, id int, body DeviceUpdate) (device Device, err error) {
	err = d.client.Update(ctx, getSpecificDevicePath(id), body, &device)
	return
}

func (d DeviceService) Delete(ctx context.Context, id int) (err error) {
	err = d.client.Delete(ctx, getSpecificDevicePath(id))
	return
}

const devicesSegment = "/v4/macbaremetal/devices"

func getDevicesPath() string {
	return devicesSegment
}

func getSpecificDevicePath(id int) string {
	return goclient.Join(devicesSegment, id)
}
