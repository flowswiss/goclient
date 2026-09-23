package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type VolumeService struct {
	client *core.Client
}

func NewVolumeService(client *core.Client) *VolumeService {
	return &VolumeService{client: client}
}

func (v VolumeService) List(ctx context.Context, cursor core.Cursor) (list common.List[Volume], err error) {
	list.Pagination, err = v.client.List(ctx, getVolumesPath(), cursor, &list.Items)
	return
}

type VolumeGetReq struct {
	ID uint `json:"-"`
}

func (v VolumeService) Get(ctx context.Context, req VolumeGetReq) (volume Volume, err error) {
	err = v.client.Get(ctx, getSpecificVolumePath(req.ID), &volume)
	return
}

type VolumeCreateReq struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	LocationID int    `json:"location_id"`
	SnapshotID *int   `json:"snapshot_id,omitempty"`
	InstanceID *int   `json:"instance_id,omitempty"`
}

func (v VolumeService) Create(ctx context.Context, req VolumeCreateReq) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumesPath(), req, &volume)
	return
}

type VolumeUpdateReq struct {
	ID uint `json:"-"`

	Name *string `json:"name,omitempty"`
}

func (v VolumeService) Update(ctx context.Context, req VolumeUpdateReq) (volume Volume, err error) {
	err = v.client.Update(ctx, getSpecificVolumePath(req.ID), req, &volume)
	return
}

type VolumeDeleteReq struct {
	ID uint `json:"-"`
}

func (v VolumeService) Delete(ctx context.Context, req VolumeDeleteReq) (err error) {
	err = v.client.Delete(ctx, getSpecificVolumePath(req.ID))
	return
}

type VolumeAttachReq struct {
	VolumeID uint `json:"-"`

	InstanceID int `json:"instance_id"`
}

func (v VolumeService) Attach(ctx context.Context, req VolumeAttachReq) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumeInstancesPath(req.VolumeID), req, &volume)
	return
}

type VolumeDetachReq struct {
	VolumeID   uint `json:"-"`
	InstanceID uint `json:"-"`
}

func (v VolumeService) Detach(ctx context.Context, req VolumeDetachReq) (err error) {
	err = v.client.Delete(ctx, getSpecificVolumeInstancePath(req.VolumeID, req.InstanceID))
	return
}

type VolumeRevertReq struct {
	VolumeID uint `json:"-"`

	SnapshotID int `json:"snapshot_id"`
}

func (v VolumeService) Revert(ctx context.Context, req VolumeRevertReq) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumeRevertPath(req.VolumeID), req, &volume)
	return
}

type VolumeExpandReq struct {
	VolumeID uint `json:"-"`

	Size int `json:"size"`
}

func (v VolumeService) Expand(ctx context.Context, req VolumeExpandReq) (volume Volume, err error) {
	err = v.client.Create(ctx, getVolumeUpgradePath(req.VolumeID), req, &volume)
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

func getSpecificVolumePath(volumeID uint) string {
	return core.Join(volumesSegment, volumeID)
}

func getVolumeInstancesPath(volumeID uint) string {
	return core.Join(volumesSegment, volumeID, volumeInstancesSegment)
}

func getSpecificVolumeInstancePath(volumeID, instanceID uint) string {
	return core.Join(volumesSegment, volumeID, volumeInstancesSegment, instanceID)
}

func getVolumeRevertPath(volumeID uint) string {
	return core.Join(volumesSegment, volumeID, volumeRevertSegment)
}

func getVolumeUpgradePath(volumeID uint) string {
	return core.Join(volumesSegment, volumeID, volumeUpgradeSegment)
}
