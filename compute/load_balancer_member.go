package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type LoadBalancerMemberService struct {
	client *core.Client
}

func NewLoadBalancerMemberService(client *core.Client) *LoadBalancerMemberService {
	return &LoadBalancerMemberService{client: client}
}

type LoadBalancerMemberListReq struct {
	LoadBalancerID     uint `json:"-"`
	LoadBalancerPoolID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (m LoadBalancerMemberService) List(ctx context.Context, req LoadBalancerMemberListReq) (
	list common.List[LoadBalancerMember],
	err error,
) {
	list.Pagination, err = m.client.List(ctx, getLoadBalancerMembersPath(req.LoadBalancerID, req.LoadBalancerPoolID), req.Cursor, &list.Items)
	return
}

type LoadBalancerMemberCreateReq struct {
	LoadBalancerID     uint `json:"-"`
	LoadBalancerPoolID uint `json:"-"`

	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (m LoadBalancerMemberService) Create(
	ctx context.Context,
	req LoadBalancerMemberCreateReq,
) (member LoadBalancerMember, err error) {
	err = m.client.Create(ctx, getLoadBalancerMembersPath(req.LoadBalancerID, req.LoadBalancerPoolID), req, &member)
	return
}

type LoadBalancerMemberDeleteReq struct {
	LoadBalancerID       uint `json:"-"`
	LoadBalancerPoolID   uint `json:"-"`
	LoadBalancerMemberID uint `json:"-"`
}

func (m LoadBalancerMemberService) Delete(ctx context.Context, req LoadBalancerMemberDeleteReq) (err error) {
	err = m.client.Delete(ctx, getSpecificLoadBalancerMemberPath(req.LoadBalancerID, req.LoadBalancerPoolID, req.LoadBalancerMemberID))
	return
}

const loadBalancerMembersSegment = "members"

func getLoadBalancerMembersPath(loadBalancerID, poolID uint) string {
	return core.Join(loadBalancerSegment, loadBalancerID, loadBalancerPoolsSegment, poolID, loadBalancerMembersSegment)
}

func getSpecificLoadBalancerMemberPath(loadBalancerID, poolID, memberID uint) string {
	return core.Join(loadBalancerSegment, loadBalancerID, loadBalancerPoolsSegment, poolID, loadBalancerMembersSegment, memberID)
}
