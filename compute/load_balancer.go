package compute

import (
	"context"
	"time"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type LoadBalancer struct {
	ID       int                             `json:"id"`
	Name     string                          `json:"name"`
	Location common.Location                 `json:"location"`
	Product  common.Product                  `json:"product"`
	Status   LoadBalancerStatus              `json:"status"`
	Networks []LoadBalancerNetworkAttachment `json:"networks"`
}

type LoadBalancerList struct {
	Items      []LoadBalancer
	Pagination goclient.Pagination
}

type LoadBalancerNetworkAttachment struct {
	Network
	Interfaces []AttachedLoadBalancerInterface `json:"network_interfaces"`
}

type AttachedLoadBalancerInterface struct {
	ID        int    `json:"id"`
	PrivateIP string `json:"private_ip"`
	PublicIP  string `json:"public_ip"`
}

type LoadBalancerCreate struct {
	Name             string `json:"name"`
	LocationID       int    `json:"location_id"`
	AttachExternalIP bool   `json:"attach_external_ip"`
	NetworkID        int    `json:"network_id"`
	PrivateIP        string `json:"private_ip"`
}

type LoadBalancerUpdate struct {
	Name string `json:"name,omitempty"`
}

type LoadBalancerPerform struct {
	Action string `json:"action"`
}

type LoadBalancerService struct {
	client goclient.Client
}

func NewLoadBalancerService(client goclient.Client) LoadBalancerService {
	return LoadBalancerService{client: client}
}

func (l LoadBalancerService) Pools(loadBalancerID int) LoadBalancerPoolService {
	return NewLoadBalancerPoolService(l.client, loadBalancerID)
}

func (l LoadBalancerService) List(ctx context.Context, cursor goclient.Cursor) (list LoadBalancerList, err error) {
	list.Pagination, err = l.client.List(ctx, getLoadBalancersPath(), cursor, &list.Items)
	return
}

func (l LoadBalancerService) Get(ctx context.Context, id int) (loadBalancer LoadBalancer, err error) {
	err = l.client.Get(ctx, getSpecificLoadBalancerPath(id), &loadBalancer)
	return
}

func (l LoadBalancerService) Create(ctx context.Context, body LoadBalancerCreate) (ordering common.Ordering, err error) {
	err = l.client.Create(ctx, getLoadBalancersPath(), body, &ordering)
	return
}

func (l LoadBalancerService) Perform(ctx context.Context, id int, body LoadBalancerPerform) (loadBalancer LoadBalancer, err error) {
	err = l.client.Create(ctx, getLoadBalancerActionPath(id), body, &loadBalancer)
	return
}

func (l LoadBalancerService) Update(ctx context.Context, id int, body LoadBalancerUpdate) (loadBalancer LoadBalancer, err error) {
	err = l.client.Update(ctx, getSpecificLoadBalancerPath(id), body, &loadBalancer)
	return
}

func (l LoadBalancerService) Delete(ctx context.Context, id int) (err error) {
	err = l.client.Delete(ctx, getSpecificLoadBalancerPath(id))
	return
}

func (l LoadBalancerService) WaitUntilMutable(ctx context.Context, id int) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			loadBalancer, err := l.Get(ctx, id)
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

func getSpecificLoadBalancerPath(loadBalancerID int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerID)
}

func getLoadBalancerActionPath(loadBalancerID int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerID, loadBalancerActionSegment)
}
