package compute

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

const (
	VolumeStatusAvailable = iota + 1
	VolumeStatusInUse
	VolumeStatusWorking
	VolumeStatusError
)

type VolumeStatus struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Volume struct {
	Id           int             `json:"id"`
	Product      common.Product  `json:"product"`
	Location     common.Location `json:"location"`
	Status       VolumeStatus    `json:"status"`
	Name         string          `json:"name"`
	Size         int             `json:"size"`
	SerialNumber string          `json:"serial"`
	Snapshots    int             `json:"snapshots"`
	Bootable     bool            `json:"bootable"`
	RootVolume   bool            `json:"root_volume"`
	AttachedTo   Server          `json:"instance"`
	CreatedAt    common.Time     `json:"created_at"`
}

type VolumeList struct {
	Items      []Volume
	Pagination goclient.Pagination
}

type VolumeCreate struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	LocationId int    `json:"location_id"`
	SnapshotId int    `json:"snapshot_id,omitempty"`
	InstanceId int    `json:"instance_id,omitempty"`
}

type VolumeUpdate struct {
	Name string `json:"name"`
}

type VolumeAttach struct {
	InstanceId int `json:"instance_id"`
}

type VolumeRevert struct {
	SnapshotId int `json:"snapshot_id"`
}

type VolumeExpand struct {
	Size int `json:"size"`
}

type VolumeService struct {
	client goclient.Client
}

func NewVolumeService(client goclient.Client) VolumeService {
	return VolumeService{client: client}
}

func (v VolumeService) List(ctx context.Context, cursor goclient.Cursor) (list VolumeList, err error) {
	list.Pagination, err = v.client.List(ctx, getVolumesPath(), cursor, &list.Items)
	return
}

func (v VolumeService) Get(ctx context.Context, id int) (volume Volume, err error) {
	err = v.client.Get(ctx, getSpecificVolumePath(id), &volume)
	return
}

func (v VolumeService) Create(ctx context.Context, body VolumeCreate) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumesPath(), body, &volume)
	return
}

func (v VolumeService) Update(ctx context.Context, id int, body VolumeUpdate) (volume Volume, err error) {
	err = v.client.Update(ctx, getSpecificVolumePath(id), body, &volume)
	return
}

func (v VolumeService) Delete(ctx context.Context, id int) (err error) {
	err = v.client.Delete(ctx, getSpecificVolumePath(id))
	return
}

func (v VolumeService) Attach(ctx context.Context, id int, body VolumeAttach) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumeInstancesPath(id), body, &volume)
	return
}

func (v VolumeService) Detach(ctx context.Context, id int, instanceId int) (err error) {
	err = v.client.Delete(ctx, getSpecificVolumeInstancePath(id, instanceId))
	return
}

func (v VolumeService) Revert(ctx context.Context, id int, body VolumeRevert) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumeRevertPath(id), body, &volume)
	return
}

func (v VolumeService) Expand(ctx context.Context, id int, body VolumeExpand) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumeUpgradePath(id), body, &volume)
	return
}

const (
	volumesSegment         = "/v4/compute/volumes"
	volumeInstancesSegment = "instances"
	volumeRevertSegment    = "revert"
	volumeUpgradeSegment   = "upgrade"
)

func getVolumesPath() string {
	return volumesSegment
}

func getSpecificVolumePath(volumeId int) string {
	return goclient.Join(volumesSegment, volumeId)
}

func getVolumeInstancesPath(volumeId int) string {
	return goclient.Join(volumesSegment, volumeId, volumeInstancesSegment)
}

func getSpecificVolumeInstancePath(volumeId, instanceId int) string {
	return goclient.Join(volumesSegment, volumeId, volumeInstancesSegment, instanceId)
}

func getVolumeRevertPath(volumeId int) string {
	return goclient.Join(volumesSegment, volumeId, volumeRevertSegment)
}

func getVolumeUpgradePath(volumeId int) string {
	return goclient.Join(volumesSegment, volumeId, volumeUpgradeSegment)
}
