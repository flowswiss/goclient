package objectstorage

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type Instance struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Location common.Location `json:"location"`
}

type InstanceService struct {
	client *core.Client
}

func NewInstanceService(client *core.Client) *InstanceService {
	return &InstanceService{
		client: client,
	}
}

func (i InstanceService) List(ctx context.Context, cursor core.Cursor) (list common.List[Instance], err error) {
	list.Pagination, err = i.client.List(ctx, getInstancePath(), cursor, &list.Items)
	return
}

type InstanceCreateReq struct {
	LocationID int `json:"location_id"`
}

func (i InstanceService) Create(ctx context.Context, req InstanceCreateReq) (instance Instance, err error) {
	err = i.client.Create(ctx, getInstancePath(), req, &instance)
	return
}

type InstanceDeleteReq struct {
	ID uint `json:"-"`
}

func (i InstanceService) Delete(ctx context.Context, req InstanceDeleteReq) (err error) {
	err = i.client.Delete(ctx, getSpecificInstancePath(req.ID))
	return
}

const instanceSegment = "/v4/object-storage/instances"

func getInstancePath() string {
	return instanceSegment
}

func getSpecificInstancePath(loadBalancerID uint) string {
	return core.Join(instanceSegment, loadBalancerID)
}
