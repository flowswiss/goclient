package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/core"
)

type LoadBalancerElasticIPService struct {
	client *core.Client
}

func NewLoadBalancerElasticIPService(client *core.Client) *LoadBalancerElasticIPService {
	return &LoadBalancerElasticIPService{
		client: client,
	}
}

type LoadBalancerElasticIPCreateReq struct {
	LoadBalancerID uint `json:"-"`

	ElasticIPID int `json:"elastic_ip_id"`
}

func (l LoadBalancerElasticIPService) Create(ctx context.Context, req LoadBalancerElasticIPCreateReq) (
	loadBalancer LoadBalancer,
	err error,
) {
	err = l.client.Create(ctx, getLoadBalancerElasticIPPath(req.LoadBalancerID), req, &loadBalancer)
	return
}

type LoadBalancerElasticIPDeleteReq struct {
	LoadBalancerID uint `json:"-"`
}

func (l LoadBalancerElasticIPService) Delete(ctx context.Context, req LoadBalancerElasticIPDeleteReq) (err error) {
	err = l.client.Delete(ctx, getLoadBalancerElasticIPPath(req.LoadBalancerID))
	return
}

const loadBalancerElasticIPSegment = "elastic-ip"

func getLoadBalancerElasticIPPath(loadBalancerID uint) string {
	return core.Join(loadBalancerSegment, loadBalancerID, loadBalancerElasticIPSegment)
}
