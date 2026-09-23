package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type SnapshotService struct {
	client *core.Client
}

func NewSnapshotService(client *core.Client) *SnapshotService {
	return &SnapshotService{client: client}
}

func (s SnapshotService) List(ctx context.Context, cursor core.Cursor) (list common.List[Snapshot], err error) {
	list.Pagination, err = s.client.List(ctx, getSnapshotsPath(), cursor, &list.Items)
	return
}

type SnapshotGetReq struct {
	ID uint `json:"-"`
}

func (s SnapshotService) Get(ctx context.Context, req SnapshotGetReq) (snapshot Snapshot, err error) {
	err = s.client.Get(ctx, getSpecificSnapshotPath(req.ID), &snapshot)
	return
}

type SnapshotCreateReq struct {
	Name     string `json:"name"`
	VolumeID int    `json:"volume_id"`
}

func (s SnapshotService) Create(ctx context.Context, req SnapshotCreateReq) (snapshot Snapshot, err error) {
	err = s.client.Create(ctx, getSnapshotsPath(), req, &snapshot)
	return
}

type SnapshotUpdateReq struct {
	ID uint `json:"-"`

	Name *string `json:"name,omitempty"`
}

func (s SnapshotService) Update(ctx context.Context, req SnapshotUpdateReq) (snapshot Snapshot, err error) {
	err = s.client.Update(ctx, getSpecificSnapshotPath(req.ID), req, &snapshot)
	return
}

type SnapshotDeleteReq struct {
	ID uint `json:"-"`
}

func (s SnapshotService) Delete(ctx context.Context, req SnapshotDeleteReq) (err error) {
	err = s.client.Delete(ctx, getSpecificSnapshotPath(req.ID))
	return
}

const snapshotsSegment = "/v4/compute/snapshots"

func getSnapshotsPath() string {
	return snapshotsSegment
}

func getSpecificSnapshotPath(snapshotID uint) string {
	return core.Join(snapshotsSegment, snapshotID)
}
