package compute

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Snapshot struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Size      int            `json:"size"`
	Status    SnapshotStatus `json:"status"`
	Volume    Volume         `json:"volume"`
	Product   common.Product `json:"product"`
	CreatedAt common.Time    `json:"created_at"`
}

const (
	SnapshotStatusAvailable = 1
	SnapshotStatusCreating  = 2
	SnapshotStatusError     = 3
)

type SnapshotStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type SnapshotList struct {
	Items      []Snapshot
	Pagination goclient.Pagination
}

type SnapshotCreate struct {
	Name     string `json:"name"`
	VolumeID int    `json:"volume_id"`
}

type SnapshotUpdate struct {
	Name string `json:"name,omitempty"`
}

type SnapshotService struct {
	client goclient.Client
}

func NewSnapshotService(client goclient.Client) SnapshotService {
	return SnapshotService{client: client}
}

func (s SnapshotService) List(ctx context.Context, cursor goclient.Cursor) (list SnapshotList, err error) {
	list.Pagination, err = s.client.List(ctx, getSnapshotsPath(), cursor, &list.Items)
	return
}

func (s SnapshotService) Get(ctx context.Context, id int) (snapshot Snapshot, err error) {
	err = s.client.Get(ctx, getSpecificSnapshotPath(id), &snapshot)
	return
}

func (s SnapshotService) Create(ctx context.Context, body SnapshotCreate) (snapshot Snapshot, err error) {
	err = s.client.Create(ctx, getSnapshotsPath(), body, &snapshot)
	return
}

func (s SnapshotService) Update(ctx context.Context, id int, body SnapshotUpdate) (snapshot Snapshot, err error) {
	err = s.client.Update(ctx, getSpecificSnapshotPath(id), body, &snapshot)
	return
}

func (s SnapshotService) Delete(ctx context.Context, id int) (err error) {
	err = s.client.Delete(ctx, getSpecificSnapshotPath(id))
	return
}

const snapshotsSegment = "/v4/compute/snapshots"

func getSnapshotsPath() string {
	return snapshotsSegment
}

func getSpecificSnapshotPath(snapshotID int) string {
	return goclient.Join(snapshotsSegment, snapshotID)
}
