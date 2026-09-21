package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type SecurityGroupService struct {
	client *core.Client
}

func NewSecurityGroupService(client *core.Client) *SecurityGroupService {
	return &SecurityGroupService{client: client}
}

func (s SecurityGroupService) List(ctx context.Context, cursor core.Cursor) (
	list common.List[SecurityGroup],
	err error,
) {
	list.Pagination, err = s.client.List(ctx, getSecurityGroupsPath(), cursor, &list.Items)
	return
}

type SecurityGroupCreateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	NetworkID   int    `json:"network_id"`
}

func (s SecurityGroupService) Create(ctx context.Context, req SecurityGroupCreateReq) (
	securityGroup SecurityGroup,
	err error,
) {
	err = s.client.Create(ctx, getSecurityGroupsPath(), req, &securityGroup)
	return
}

type SecurityGroupGetReq struct {
	ID uint `json:"-"`
}

func (s SecurityGroupService) Get(ctx context.Context, req SecurityGroupGetReq) (
	securityGroup SecurityGroup,
	err error,
) {
	err = s.client.Get(ctx, getSpecificSecurityGroupPath(req.ID), &securityGroup)
	return
}

type SecurityGroupUpdateReq struct {
	ID uint `json:"-"`

	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (s SecurityGroupService) Update(ctx context.Context, req SecurityGroupUpdateReq) (
	securityGroup SecurityGroup,
	err error,
) {
	err = s.client.Update(ctx, getSpecificSecurityGroupPath(req.ID), req, &securityGroup)
	return
}

type SecurityGroupDeleteReq struct {
	ID uint `json:"-"`
}

func (s SecurityGroupService) Delete(ctx context.Context, req SecurityGroupDeleteReq) (err error) {
	err = s.client.Delete(ctx, getSpecificSecurityGroupPath(req.ID))
	return
}

const securityGroupsSegment = "/v4/macbaremetal/security-groups"

func getSecurityGroupsPath() string {
	return securityGroupsSegment
}

func getSpecificSecurityGroupPath(id uint) string {
	return core.Join(securityGroupsSegment, id)
}
