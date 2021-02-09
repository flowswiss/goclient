package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type LoadBalancerPool struct {
	Id             int                     `json:"id"`
	Name           string                  `json:"name"`
	Status         LoadBalancerStatus      `json:"status"`
	EntryProtocol  LoadBalancerProtocol    `json:"entry_protocol"`
	TargetProtocol LoadBalancerProtocol    `json:"target_protocol"`
	EntryPort      int                     `json:"entry_port"`
	Algorithm      LoadBalancerAlgorithm   `json:"algorithm"`
	StickySession  bool                    `json:"sticky_session"`
	HealthCheck    LoadBalancerHealthCheck `json:"health_check"`
}

type LoadBalancerPoolList struct {
	Items      []LoadBalancerPool
	Pagination goclient.Pagination
}

type LoadBalancerHealthCheck struct {
	Type               LoadBalancerHealthCheckType `json:"type"`
	HttpMethod         string                      `json:"http_method"`
	HttpPath           string                      `json:"http_path"`
	Interval           int                         `json:"interval"`
	Timeout            int                         `json:"timeout"`
	HealthyThreshold   int                         `json:"healthy_threshold"`
	UnhealthyThreshold int                         `json:"unhealthy_threshold"`
}

type LoadBalancerPoolCreate struct {
	EntryProtocolId      int                            `json:"entry_protocol_id"`
	TargetProtocolId     int                            `json:"target_protocol_id"`
	CertificateId        int                            `json:"certificate_id,omitempty"`
	EntryPort            int                            `json:"entry_port"`
	BalancingAlgorithmId int                            `json:"balancing_algorithm_id"`
	StickySession        bool                           `json:"sticky_session"`
	Members              []LoadBalancerMemberCreate     `json:"members,omitempty"`
	HealthCheck          LoadBalancerHealthCheckOptions `json:"health_check"`
}

type LoadBalancerPoolUpdate struct {
	CertificateId        int                            `json:"certificate_id,omitempty"`
	BalancingAlgorithmId int                            `json:"balancing_algorithm_id,omitempty"`
	StickySession        bool                           `json:"sticky_session,omitempty"`
	HealthCheck          LoadBalancerHealthCheckOptions `json:"health_check,omitempty"`
}

type LoadBalancerHealthCheckOptions struct {
	TypeId             int    `json:"type_id"`
	HttpMethod         string `json:"http_method,omitempty"`
	HttpPath           string `json:"http_path,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	HealthyThreshold   int    `json:"healthy_threshold,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
}

type LoadBalancerPoolService struct {
	client goclient.Client

	loadBalancerId int
}

func NewLoadBalancerPoolService(client goclient.Client, loadBalancerId int) LoadBalancerPoolService {
	return LoadBalancerPoolService{client: client, loadBalancerId: loadBalancerId}
}

func (l LoadBalancerPoolService) Members(poolId int) LoadBalancerMemberService {
	return NewLoadBalancerMemberService(l.client, l.loadBalancerId, poolId)
}

func (l LoadBalancerPoolService) List(ctx context.Context, cursor goclient.Cursor) (list LoadBalancerPoolList, err error) {
	list.Pagination, err = l.client.List(ctx, getLoadBalancerPoolsPath(l.loadBalancerId), cursor, &list.Items)
	return
}

func (l LoadBalancerPoolService) Get(ctx context.Context, id int) (pool LoadBalancerPool, err error) {
	err = l.client.Get(ctx, getSpecificLoadBalancerPoolPath(l.loadBalancerId, id), &pool)
	return
}

func (l LoadBalancerPoolService) Create(ctx context.Context, body LoadBalancerPoolCreate) (pool LoadBalancerPool, err error) {
	err = l.client.Create(ctx, getLoadBalancerPoolsPath(l.loadBalancerId), body, &pool)
	return
}

func (l LoadBalancerPoolService) Update(ctx context.Context, id int, body LoadBalancerPoolUpdate) (pool LoadBalancerPool, err error) {
	err = l.client.Create(ctx, getSpecificLoadBalancerPoolPath(l.loadBalancerId, id), body, &pool)
	return
}

func (l LoadBalancerPoolService) Delete(ctx context.Context, id int) (err error) {
	err = l.client.Delete(ctx, getSpecificLoadBalancerPoolPath(l.loadBalancerId, id))
	return
}

const loadBalancerPoolsSegment = "balancing-pools"

func getLoadBalancerPoolsPath(loadBalancerId int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerId, loadBalancerPoolsSegment)
}

func getSpecificLoadBalancerPoolPath(loadBalancerId, poolId int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerId, loadBalancerPoolsSegment, poolId)
}
