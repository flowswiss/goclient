package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

const (
	DirectionIngress = "ingress"
	DirectionEgress  = "egress"

	ProtocolAny  = -1
	ProtocolICMP = 1
	ProtocolTCP  = 6
	ProtocolUDP  = 17
)

type SecurityGroupRule struct {
	ID                  int           `json:"id"`
	Direction           string        `json:"direction"`
	Protocol            int           `json:"protocol"`
	FromPort            int           `json:"from_port"`
	ToPort              int           `json:"to_port"`
	ICMPType            int           `json:"icmp_type"`
	ICMPCode            int           `json:"icmp_code"`
	IPRange             string        `json:"ip_range"`
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
	ICMPType              int    `json:"icmp_type,omitempty"`
	ICMPCode              int    `json:"icmp_code,omitempty"`
	IPRange               string `json:"ip_range,omitempty"`
	RemoteSecurityGroupID int    `json:"remote_security_group_id,omitempty"`
}

type SecurityGroupRuleService struct {
	client          goclient.Client
	securityGroupID int
}

func NewSecurityGroupRuleService(client goclient.Client, securityGroupID int) SecurityGroupRuleService {
	return SecurityGroupRuleService{client: client, securityGroupID: securityGroupID}
}

func (s SecurityGroupRuleService) List(ctx context.Context, cursor goclient.Cursor) (list SecurityGroupRuleList, err error) {
	list.Pagination, err = s.client.List(ctx, getSecurityGroupRulesPath(s.securityGroupID), cursor, &list.Items)
	return
}

func (s SecurityGroupRuleService) Create(ctx context.Context, body SecurityGroupRuleOptions) (rule SecurityGroupRule, err error) {
	err = s.client.Create(ctx, getSecurityGroupRulesPath(s.securityGroupID), body, &rule)
	return
}

func (s SecurityGroupRuleService) Update(ctx context.Context, id int, body SecurityGroupRuleOptions) (rule SecurityGroupRule, err error) {
	err = s.client.Update(ctx, getSpecificSecurityGroupRulePath(s.securityGroupID, id), body, &rule)
	return
}

func (s SecurityGroupRuleService) Delete(ctx context.Context, id int) (err error) {
	err = s.client.Delete(ctx, getSpecificSecurityGroupRulePath(s.securityGroupID, id))
	return
}

const securityGroupRulesSegment = "rules"

func getSecurityGroupRulesPath(securityGroupID int) string {
	return goclient.Join(securityGroupsSegment, securityGroupID, securityGroupRulesSegment)
}

func getSpecificSecurityGroupRulePath(securityGroupID, ruleID int) string {
	return goclient.Join(securityGroupsSegment, securityGroupID, securityGroupRulesSegment, ruleID)
}
