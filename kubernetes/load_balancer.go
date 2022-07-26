package kubernetes

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/compute"
)

type LoadBalancer = compute.LoadBalancer
type LoadBalancerList = compute.LoadBalancerList

type LoadBalancerService struct {
	client    goclient.Client
	clusterID int
}

func NewLoadBalancerService(client goclient.Client, clusterID int) LoadBalancerService {
	return LoadBalancerService{
		client:    client,
		clusterID: clusterID,
	}
}

func (v LoadBalancerService) List(ctx context.Context, cursor goclient.Cursor) (list LoadBalancerList, err error) {
	list.Pagination, err = v.client.List(ctx, getLoadBalancerPath(v.clusterID), cursor, &list.Items)
	return
}

const loadBalancerSegment = "load-balancers"

func getLoadBalancerPath(clusterID int) string {
	return goclient.Join(clusterSegment, clusterID, loadBalancerSegment)
}
