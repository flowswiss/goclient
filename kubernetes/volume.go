package kubernetes

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/compute"
	"github.com/flowswiss/goclient/v2/core"
)

type Volume = compute.Volume

type VolumeService struct {
	client *core.Client
}

func NewVolumeService(client *core.Client) *VolumeService {
	return &VolumeService{
		client: client,
	}
}

type VolumeListReq struct {
	ClusterID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (v VolumeService) List(ctx context.Context, req VolumeListReq) (list common.List[Volume], err error) {
	list.Pagination, err = v.client.List(ctx, getVolumePath(req.ClusterID), req.Cursor, &list.Items)
	return
}

type VolumeDeleteReq struct {
	ClusterID uint `json:"-"`
	VolumeID  uint `json:"-"`
}

func (v VolumeService) Delete(ctx context.Context, req VolumeDeleteReq) (err error) {
	err = v.client.Delete(ctx, getSpecificVolumePath(req.ClusterID, req.VolumeID))
	return
}

const volumeSegment = "volumes"

func getVolumePath(clusterID uint) string {
	return core.Join(clusterSegment, clusterID, volumeSegment)
}

func getSpecificVolumePath(clusterID, volumeID uint) string {
	return core.Join(clusterSegment, clusterID, volumeSegment, volumeID)
}
