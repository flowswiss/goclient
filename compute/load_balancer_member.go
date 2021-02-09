package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type LoadBalancerMember struct {
	Id      int                `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Port    int                `json:"port"`
	Status  LoadBalancerStatus `json:"status"`
}

type LoadBalancerMemberList struct {
	Items      []LoadBalancerMember
	Pagination goclient.Pagination
}

type LoadBalancerMemberCreate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type LoadBalancerMemberService struct {
	client goclient.Client

	loadBalancerId int
	poolId         int
}

func NewLoadBalancerMemberService(client goclient.Client, loadBalancerId, poolId int) LoadBalancerMemberService {
	return LoadBalancerMemberService{client: client, loadBalancerId: loadBalancerId, poolId: poolId}
}

func (m LoadBalancerMemberService) List(ctx context.Context, cursor goclient.Cursor) (list LoadBalancerMemberList, err error) {
	list.Pagination, err = m.client.List(ctx, getLoadBalancerMembersPath(m.loadBalancerId, m.poolId), cursor, &list.Items)
	return
}

func (m LoadBalancerMemberService) Create(ctx context.Context, body LoadBalancerMemberCreate) (member LoadBalancerMember, err error) {
	err = m.client.Create(ctx, getLoadBalancerMembersPath(m.loadBalancerId, m.poolId), body, &member)
	return
}

func (m LoadBalancerMemberService) Delete(ctx context.Context, id int) (err error) {
	err = m.client.Delete(ctx, getSpecificLoadBalancerMemberPath(m.loadBalancerId, m.poolId, id))
	return
}

const loadBalancerMembersSegment = "members"

func getLoadBalancerMembersPath(loadBalancerId, poolId int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerId, loadBalancerPoolsSegment, poolId, loadBalancerMembersSegment)
}

func getSpecificLoadBalancerMemberPath(loadBalancerId, poolId, memberId int) string {
	return goclient.Join(loadBalancerSegment, loadBalancerId, loadBalancerPoolsSegment, poolId, loadBalancerMembersSegment, memberId)
}
