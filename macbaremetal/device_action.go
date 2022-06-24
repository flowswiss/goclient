package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient"
)

type DeviceAction struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}

type DeviceStatus struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Key     string         `json:"key"`
	Actions []DeviceAction `json:"actions"`
}

type DeviceRunAction struct {
	Action string `json:"action"`
}

type DeviceWorkflow struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}

type DeviceWorkflowList struct {
	Items      []DeviceWorkflow
	Pagination goclient.Pagination
}

type DeviceRunWorkflow struct {
	Workflow string `json:"workflow"`
}

type DeviceActionService struct {
	client   goclient.Client
	deviceID int
}

func NewDeviceActionService(client goclient.Client, deviceID int) DeviceActionService {
	return DeviceActionService{client: client, deviceID: deviceID}
}

func (d DeviceActionService) Run(ctx context.Context, body DeviceRunAction) (device Device, err error) {
	err = d.client.Create(ctx, getDeviceActionPath(d.deviceID), body, &device)
	return
}

type DeviceWorkflowService struct {
	client   goclient.Client
	deviceID int
}

func NewDeviceWorkflowService(client goclient.Client, deviceID int) DeviceWorkflowService {
	return DeviceWorkflowService{client: client, deviceID: deviceID}
}

func (d DeviceWorkflowService) List(ctx context.Context, cursor goclient.Cursor) (list DeviceWorkflowList, err error) {
	list.Pagination, err = d.client.List(ctx, getDeviceWorkflowPath(d.deviceID), cursor, &list.Items)
	return
}

func (d DeviceWorkflowService) Run(ctx context.Context, body DeviceRunWorkflow) (device Device, err error) {
	err = d.client.Create(ctx, getDeviceWorkflowPath(d.deviceID), body, &device)
	return
}

const (
	deviceActionSegment   = "actions"
	deviceWorkflowSegment = "workflows"
)

func getDeviceActionPath(id int) string {
	return goclient.Join(devicesSegment, id, deviceActionSegment)
}

func getDeviceWorkflowPath(id int) string {
	return goclient.Join(devicesSegment, id, deviceWorkflowSegment)
}
