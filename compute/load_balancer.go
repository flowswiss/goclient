package compute

import (
	"context"
	"time"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type LoadBalancerService struct {
	client *core.Client
}

func NewLoadBalancerService(client *core.Client) *LoadBalancerService {
	return &LoadBalancerService{client: client}
}

func (l LoadBalancerService) List(ctx context.Context, cursor core.Cursor) (list common.List[LoadBalancer], err error) {
	list.Pagination, err = l.client.List(ctx, getLoadBalancersPath(), cursor, &list.Items)
	return
}

type LoadBalancerGetReq struct {
	ID uint `json:"-"`
}

func (l LoadBalancerService) Get(ctx context.Context, req LoadBalancerGetReq) (loadBalancer LoadBalancer, err error) {
	err = l.client.Get(ctx, getSpecificLoadBalancerPath(req.ID), &loadBalancer)
	return
}

type LoadBalancerCreateReq struct {
	Name             string `json:"name"`
	LocationID       int    `json:"location_id"`
	AttachExternalIP bool   `json:"attach_external_ip"`
	NetworkID        int    `json:"network_id"`
	PrivateIP        string `json:"private_ip"`
}

func (l LoadBalancerService) Create(ctx context.Context, req LoadBalancerCreateReq) (
	ordering common.Ordering,
	err error,
) {
	err = l.client.Create(ctx, getLoadBalancersPath(), req, &ordering)
	return
}

type LoadBalancerPerformReq struct {
	ID uint `json:"-"`

	Action string `json:"action"`
}

func (l LoadBalancerService) Perform(ctx context.Context, req LoadBalancerPerformReq) (
	loadBalancer LoadBalancer,
	err error,
) {
	err = l.client.Create(ctx, getLoadBalancerActionPath(req.ID), req, &loadBalancer)
	return
}

type LoadBalancerUpdateReq struct {
	ID uint `json:"-"`

	Name string `json:"name,omitempty"`
}

func (l LoadBalancerService) Update(ctx context.Context, req LoadBalancerUpdateReq) (
	loadBalancer LoadBalancer,
	err error,
) {
	err = l.client.Update(ctx, getSpecificLoadBalancerPath(req.ID), req, &loadBalancer)
	return
}

type LoadBalancerDeleteReq struct {
	ID uint `json:"-"`
}

func (l LoadBalancerService) Delete(ctx context.Context, req LoadBalancerDeleteReq) (err error) {
	err = l.client.Delete(ctx, getSpecificLoadBalancerPath(req.ID))
	return
}

func (l LoadBalancerService) WaitUntilMutable(ctx context.Context, req LoadBalancerGetReq) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			loadBalancer, err := l.Get(ctx, req)
			if err != nil {
				return err
			}

			if loadBalancer.Status.ID != LoadBalancerStatusWorking {
				return nil
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

const (
	loadBalancerSegment       = "/v4/compute/load-balancers"
	loadBalancerActionSegment = "action"
)

func getLoadBalancersPath() string {
	return loadBalancerSegment
}

func getSpecificLoadBalancerPath(loadBalancerID uint) string {
	return core.Join(loadBalancerSegment, loadBalancerID)
}

func getLoadBalancerActionPath(loadBalancerID uint) string {
	return core.Join(loadBalancerSegment, loadBalancerID, loadBalancerActionSegment)
}
