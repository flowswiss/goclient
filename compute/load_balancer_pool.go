package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type LoadBalancerPool struct {
	ID             int                     `json:"id"`
	Name           string                  `json:"name"`
	Status         LoadBalancerStatus      `json:"status"`
	EntryProtocol  LoadBalancerProtocol    `json:"entry_protocol"`
	TargetProtocol LoadBalancerProtocol    `json:"target_protocol"`
	Certificate    Certificate             `json:"certificate"`
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
	HTTPMethod         string                      `json:"http_method"`
	HTTPPath           string                      `json:"http_path"`
	Interval           int                         `json:"interval"`
	Timeout            int                         `json:"timeout"`
	HealthyThreshold   int                         `json:"healthy_threshold"`
	UnhealthyThreshold int                         `json:"unhealthy_threshold"`
}

type LoadBalancerPoolCreate struct {
	EntryProtocolID      int                            `json:"entry_protocol_id"`
	TargetProtocolID     int                            `json:"target_protocol_id"`
	CertificateID        int                            `json:"certificate_id,omitempty"`
	EntryPort            int                            `json:"entry_port"`
	BalancingAlgorithmID int                            `json:"balancing_algorithm_id"`
	StickySession        bool                           `json:"sticky_session"`
	Members              []LoadBalancerMemberCreate     `json:"members,omitempty"`
	HealthCheck          LoadBalancerHealthCheckOptions `json:"health_check"`
}

type LoadBalancerPoolUpdate struct {
	CertificateID        int                            `json:"certificate_id,omitempty"`
	BalancingAlgorithmID int                            `json:"balancing_algorithm_id,omitempty"`
	StickySession        bool                           `json:"sticky_session,omitempty"`
	HealthCheck          LoadBalancerHealthCheckOptions `json:"health_check,omitempty"`
}

type LoadBalancerHealthCheckOptions struct {
	TypeID             int    `json:"type_id"`
	HTTPMethod         string `json:"http_method,omitempty"`
	HTTPPath           string `json:"http_path,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	HealthyThreshold   int    `json:"healthy_threshold,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
}

type LoadBalancerPoolService struct {
	client goclient.Client

	loadBalancerID int
}

func NewLoadBalancerPoolService(client goclient.Client, loadBalancerID int) LoadBalancerPoolService {
	return LoadBalancerPoolService{client: client, loadBalancerID: loadBalancerID}
}

func (l LoadBalancerPoolService) Members(poolID int) LoadBalancerMemberService {
	return NewLoadBalancerMemberService(l.client, l.loadBalancerID, poolID)
}

func (l LoadBalancerPoolService) List(ctx context.Context, cursor goclient.Cursor) (list LoadBalancerPoolList, err error) {
	list.Pagination, err = l.client.List(ctx, getLoadBalancerPoolsPath(l.loadBalancerID), cursor, &list.Items)
	return
}

func (l LoadBalancerPoolService) Get(ctx context.Context, id int) (pool LoadBalancerPool, err error) {
	err = l.client.Get(ctx, getSpecificLoadBalancerPoolPath(l.loadBalancerID, id), &pool)
	return
}

func (l LoadBalancerPoolService) Create(ctx context.Context, body LoadBalancerPoolCreate) (pool LoadBalancerPool, err error) {
	err = l.client.Create(ctx, getLoadBalancerPoolsPath(l.loadBalancerID), body, &pool)
	return
}

func (l LoadBalancerPoolService) Update(ctx context.Context, id int, body LoadBalancerPoolUpdate) (pool LoadBalancerPool, err error) {
	err = l.client.Update(ctx, getSpecificLoadBalancerPoolPath(l.loadBalancerID, id), body, &pool)
	return
}

func (l LoadBalancerPoolService) Delete(ctx context.Context, id int) (err error) {
	err = l.client.Delete(ctx, getSpecificLoadBalancerPoolPath(l.loadBalancerID, id))
	return
}

const loadBalancerPoolsSegment = "balancing-pools"

func getLoadBalancerPoolsPath(loadBalancerID int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerID, loadBalancerPoolsSegment)
}

func getSpecificLoadBalancerPoolPath(loadBalancerID, poolID int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerID, loadBalancerPoolsSegment, poolID)
}
