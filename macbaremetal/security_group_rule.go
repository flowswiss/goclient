package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type SecurityGroupRuleService struct {
	client *core.Client
}

func NewSecurityGroupRuleService(client *core.Client) *SecurityGroupRuleService {
	return &SecurityGroupRuleService{client: client}
}

type SecurityGroupRuleListReq struct {
	SecurityGroupID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (s SecurityGroupRuleService) List(ctx context.Context, req SecurityGroupRuleListReq) (
	list common.List[SecurityGroupRule],
	err error,
) {
	list.Pagination, err = s.client.List(ctx, getSecurityGroupRulesPath(req.SecurityGroupID), req.Cursor, &list.Items)
	return
}

type SecurityGroupRuleCreateReq struct {
	SecurityGroupID uint `json:"-"`

	Direction string `json:"direction,omitempty"`
	Protocol  int    `json:"protocol,omitempty"`
	FromPort  int    `json:"from_port,omitempty"`
	ToPort    int    `json:"to_port,omitempty"`
	ICMPType  int    `json:"icmp_type,omitempty"`
	ICMPCode  int    `json:"icmp_code,omitempty"`
	IPRange   string `json:"ip_range,omitempty"`
}

func (s SecurityGroupRuleService) Create(ctx context.Context, req SecurityGroupRuleCreateReq) (
	rule SecurityGroupRule,
	err error,
) {
	err = s.client.Create(ctx, getSecurityGroupRulesPath(req.SecurityGroupID), req, &rule)
	return
}

type SecurityGroupRuleUpdateReq struct {
	SecurityGroupID     uint `json:"-"`
	SecurityGroupRuleID uint `json:"-"`

	Direction *string `json:"direction,omitempty"`
	Protocol  *int    `json:"protocol,omitempty"`
	FromPort  *int    `json:"from_port,omitempty"`
	ToPort    *int    `json:"to_port,omitempty"`
	ICMPType  *int    `json:"icmp_type,omitempty"`
	ICMPCode  *int    `json:"icmp_code,omitempty"`
	IPRange   *string `json:"ip_range,omitempty"`
}

func (s SecurityGroupRuleService) Update(
	ctx context.Context,
	req SecurityGroupRuleUpdateReq,
) (rule SecurityGroupRule, err error) {
	err = s.client.Update(ctx, getSpecificSecurityGroupRulePath(req.SecurityGroupID, req.SecurityGroupRuleID), req, &rule)
	return
}

type SecurityGroupRuleDeleteReq struct {
	SecurityGroupID     uint `json:"-"`
	SecurityGroupRuleID uint `json:"-"`
}

func (s SecurityGroupRuleService) Delete(ctx context.Context, req SecurityGroupRuleDeleteReq) (err error) {
	err = s.client.Delete(ctx, getSpecificSecurityGroupRulePath(req.SecurityGroupID, req.SecurityGroupRuleID))
	return
}

const securityGroupRulesSegment = "rules"

func getSecurityGroupRulesPath(securityGroupID uint) string {
	return core.Join(securityGroupsSegment, securityGroupID, securityGroupRulesSegment)
}

func getSpecificSecurityGroupRulePath(securityGroupID, ruleID uint) string {
	return core.Join(securityGroupsSegment, securityGroupID, securityGroupRulesSegment, ruleID)
}
