package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

const (
	DirectionIngress = "ingress"
	DirectionEgress  = "egress"

	ProtocolICMP = 1
	ProtocolTCP  = 6
	ProtocolUDP  = 17
)

type SecurityGroupRule struct {
	Id                  int           `json:"id"`
	Direction           string        `json:"direction"`
	Protocol            int           `json:"protocol"`
	FromPort            int           `json:"from_port"`
	ToPort              int           `json:"to_port"`
	IcmpType            int           `json:"icmp_type"`
	IcmpCode            int           `json:"icmp_code"`
	IpRange             string        `json:"ip_range"`
	RemoteSecurityGroup SecurityGroup `json:"remote_security_group"`
}

type SecurityGroupRuleList struct {
	Items      []SecurityGroupRule
	Pagination goclient.Pagination
}

type SecurityGroupRuleOptions struct {
	Direction             string `json:"direction"`
	Protocol              int    `json:"protocol"`
	FromPort              int    `json:"from_port,omitempty"`
	ToPort                int    `json:"to_port,omitempty"`
	IcmpType              int    `json:"icmp_type,omitempty"`
	IcmpCode              int    `json:"icmp_code,omitempty"`
	IpRange               string `json:"ip_range,omitempty"`
	RemoteSecurityGroupId int    `json:"remote_security_group_id,omitempty"`
}

type SecurityGroupRuleService struct {
	client          goclient.Client
	securityGroupId int
}

func NewSecurityGroupRuleService(client goclient.Client, securityGroupId int) SecurityGroupRuleService {
	return SecurityGroupRuleService{client: client, securityGroupId: securityGroupId}
}

func (s SecurityGroupRuleService) List(ctx context.Context, cursor goclient.Cursor) (list SecurityGroupRuleList, err error) {
	list.Pagination, err = s.client.List(ctx, getSecurityGroupRulesPath(s.securityGroupId), cursor, &list.Items)
	return
}

func (s SecurityGroupRuleService) Create(ctx context.Context, body SecurityGroupRuleOptions) (rule SecurityGroupRule, err error) {
	err = s.client.Create(ctx, getSecurityGroupRulesPath(s.securityGroupId), body, &rule)
	return
}

func (s SecurityGroupRuleService) Update(ctx context.Context, id int, body SecurityGroupRuleOptions) (rule SecurityGroupRule, err error) {
	err = s.client.Update(ctx, getSpecificSecurityGroupRulePath(s.securityGroupId, id), body, &rule)
	return
}

func (s SecurityGroupRuleService) Delete(ctx context.Context, id int) (err error) {
	err = s.client.Delete(ctx, getSpecificSecurityGroupRulePath(s.securityGroupId, id))
	return
}

const securityGroupRulesSegment = "rules"

func getSecurityGroupRulesPath(securityGroupId int) string {
	return goclient.Join(securityGroupRulesSegment, securityGroupId, securityGroupRulesSegment)
}

func getSpecificSecurityGroupRulePath(securityGroupId, ruleId int) string {
	return goclient.Join(securityGroupRulesSegment, securityGroupId, securityGroupRulesSegment, ruleId)
}
