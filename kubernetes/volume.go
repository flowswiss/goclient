package kubernetes

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/compute"
)

type Volume = compute.Volume
type VolumeList = compute.VolumeList

type VolumeService struct {
	client    goclient.Client
	clusterID int
}

func NewVolumeService(client goclient.Client, clusterID int) VolumeService {
	return VolumeService{
		client:    client,
		clusterID: clusterID,
	}
}

func (v VolumeService) List(ctx context.Context, cursor goclient.Cursor) (list VolumeList, err error) {
	list.Pagination, err = v.client.List(ctx, getVolumePath(v.clusterID), cursor, &list.Items)
	return
}

func (v VolumeService) Delete(ctx context.Context, id int) (err error) {
	err = v.client.Delete(ctx, getSpecificVolumePath(v.clusterID, id))
	return
}

const volumeSegment = "volumes"

func getVolumePath(clusterID int) string {
	return goclient.Join(clusterSegment, clusterID, volumeSegment)
}

func getSpecificVolumePath(clusterID, volumeID int) string {
	return goclient.Join(clusterSegment, clusterID, volumeSegment, volumeID)
}
