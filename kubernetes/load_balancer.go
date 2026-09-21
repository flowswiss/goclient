package kubernetes

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/compute"
	"github.com/flowswiss/goclient/v2/core"
)

type LoadBalancer = compute.LoadBalancer

type LoadBalancerService struct {
	client *core.Client
}

func NewLoadBalancerService(client *core.Client) *LoadBalancerService {
	return &LoadBalancerService{
		client: client,
	}
}

type LoadBalancerListReq struct {
	ClusterID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (v LoadBalancerService) List(ctx context.Context, req LoadBalancerListReq) (
	list common.List[LoadBalancer],
	err error,
) {
	list.Pagination, err = v.client.List(ctx, getLoadBalancerPath(req.ClusterID), req.Cursor, &list.Items)
	return
}

const loadBalancerSegment = "load-balancers"

func getLoadBalancerPath(clusterID uint) string {
	return core.Join(clusterSegment, clusterID, loadBalancerSegment)
}
