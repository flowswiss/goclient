package kubernetes

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/compute"
	"github.com/flowswiss/goclient/v2/core"
)

type Snapshot = compute.Snapshot

type SnapshotService struct {
	client *core.Client
}

func NewSnapshotService(client *core.Client) *SnapshotService {
	return &SnapshotService{
		client: client,
	}
}

type SnapshotListReq struct {
	ClusterID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (s SnapshotService) List(ctx context.Context, req SnapshotListReq) (list common.List[Snapshot], err error) {
	list.Pagination, err = s.client.List(ctx, getSnapshotsPath(req.ClusterID), req.Cursor, &list.Items)
	return
}

const snapshotsSegment = "snapshots"

func getSnapshotsPath(clusterID uint) string {
	return core.Join(clustersSegment, clusterID, snapshotsSegment)
}
