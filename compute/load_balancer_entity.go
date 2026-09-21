package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type LoadBalancerEntityService struct {
	client *core.Client
}

func NewLoadBalancerEntityService(client *core.Client) *LoadBalancerEntityService {
	return &LoadBalancerEntityService{client: client}
}

func (l LoadBalancerEntityService) ListAlgorithms(
	ctx context.Context,
	cursor core.Cursor,
) (list common.List[LoadBalancerAlgorithm], err error) {
	list.Pagination, err = l.client.List(ctx, "/v4/entities/compute/load-balancer-algorithms", cursor, &list.Items)
	return
}

func (l LoadBalancerEntityService) ListProtocols(
	ctx context.Context,
	cursor core.Cursor,
) (list common.List[LoadBalancerProtocol], err error) {
	list.Pagination, err = l.client.List(ctx, "/v4/entities/compute/load-balancer-protocols", cursor, &list.Items)
	return
}

func (l LoadBalancerEntityService) ListHealthCheckTypes(
	ctx context.Context,
	cursor core.Cursor,
) (list common.List[LoadBalancerHealthCheckType], err error) {
	list.Pagination, err = l.client.List(ctx, "/v4/entities/compute/load-balancer-health-check-types", cursor, &list.Items)
	return
}
