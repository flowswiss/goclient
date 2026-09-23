package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type LoadBalancerHealthCheckOptions struct {
	TypeID             int     `json:"type_id"`
	HTTPMethod         *string `json:"http_method,omitempty"`
	HTTPPath           *string `json:"http_path,omitempty"`
	Interval           *int    `json:"interval,omitempty"`
	Timeout            *int    `json:"timeout,omitempty"`
	HealthyThreshold   *int    `json:"healthy_threshold,omitempty"`
	UnhealthyThreshold *int    `json:"unhealthy_threshold,omitempty"`
}

type LoadBalancerPoolService struct {
	client *core.Client
}

func NewLoadBalancerPoolService(client *core.Client) *LoadBalancerPoolService {
	return &LoadBalancerPoolService{client: client}
}

type LoadBalancerPoolListReq struct {
	LoadBalancerID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (l LoadBalancerPoolService) List(ctx context.Context, req LoadBalancerPoolListReq) (
	list common.List[LoadBalancerPool],
	err error,
) {
	list.Pagination, err = l.client.List(ctx, getLoadBalancerPoolsPath(req.LoadBalancerID), req.Cursor, &list.Items)
	return
}

type LoadBalancerPoolGetReq struct {
	LoadBalancerID     uint `json:"-"`
	LoadBalancerPoolID uint `json:"-"`
}

func (l LoadBalancerPoolService) Get(ctx context.Context, req LoadBalancerPoolGetReq) (
	pool LoadBalancerPool,
	err error,
) {
	err = l.client.Get(ctx, getSpecificLoadBalancerPoolPath(req.LoadBalancerID, req.LoadBalancerPoolID), &pool)
	return
}

type LoadBalancerPoolCreateReq struct {
	LoadBalancerID uint `json:"-"`

	EntryProtocolID      int                             `json:"entry_protocol_id"`
	TargetProtocolID     int                             `json:"target_protocol_id"`
	CertificateID        *int                            `json:"certificate_id,omitempty"`
	EntryPort            int                             `json:"entry_port"`
	BalancingAlgorithmID int                             `json:"balancing_algorithm_id"`
	StickySession        bool                            `json:"sticky_session"`
	Members              []LoadBalancerMemberCreateReq   `json:"members,omitempty"`
	HealthCheck          *LoadBalancerHealthCheckOptions `json:"health_check,omitempty"`
}

func (l LoadBalancerPoolService) Create(ctx context.Context, req LoadBalancerPoolCreateReq) (
	pool LoadBalancerPool,
	err error,
) {
	err = l.client.Create(ctx, getLoadBalancerPoolsPath(req.LoadBalancerID), req, &pool)
	return
}

type LoadBalancerPoolUpdateReq struct {
	LoadBalancerID     uint `json:"-"`
	LoadBalancerPoolID uint `json:"-"`

	CertificateID        *int                            `json:"certificate_id,omitempty"`
	BalancingAlgorithmID *int                            `json:"balancing_algorithm_id,omitempty"`
	StickySession        *bool                           `json:"sticky_session,omitempty"`
	HealthCheck          *LoadBalancerHealthCheckOptions `json:"health_check,omitempty"`
}

func (l LoadBalancerPoolService) Update(
	ctx context.Context,
	req LoadBalancerPoolUpdateReq,
) (pool LoadBalancerPool, err error) {
	err = l.client.Update(ctx, getSpecificLoadBalancerPoolPath(req.LoadBalancerID, req.LoadBalancerPoolID), req, &pool)
	return
}

type LoadBalancerPoolDeleteReq struct {
	LoadBalancerID     uint `json:"-"`
	LoadBalancerPoolID uint `json:"-"`
}

func (l LoadBalancerPoolService) Delete(ctx context.Context, req LoadBalancerPoolDeleteReq) (err error) {
	err = l.client.Delete(ctx, getSpecificLoadBalancerPoolPath(req.LoadBalancerID, req.LoadBalancerPoolID))
	return
}

const loadBalancerPoolsSegment = "balancing-pools"

func getLoadBalancerPoolsPath(loadBalancerID uint) string {
	return core.Join(loadBalancerSegment, loadBalancerID, loadBalancerPoolsSegment)
}

func getSpecificLoadBalancerPoolPath(loadBalancerID, poolID uint) string {
	return core.Join(loadBalancerSegment, loadBalancerID, loadBalancerPoolsSegment, poolID)
}
