package macbaremetal

import "github.com/flowswiss/goclient/v2/core"

type Service struct {
	Device              *DeviceService
	ElasticIP           *ElasticIPService
	ElasticIPAttachment *ElasticIPAttachmentService
	Network             *NetworkService
	NetworkInterface    *NetworkInterfaceService
	Router              *RouterService
	RouterInterface     *RouterInterfaceService
	SecurityGroup       *SecurityGroupService
	SecurityGroupRule   *SecurityGroupRuleService
}

func NewService(client *core.Client) *Service {
	return &Service{
		Device:              NewDeviceService(client),
		ElasticIP:           NewElasticIPService(client),
		ElasticIPAttachment: NewElasticIPAttachmentService(client),
		Network:             NewNetworkService(client),
		NetworkInterface:    NewNetworkInterfaceService(client),
		Router:              NewRouterService(client),
		RouterInterface:     NewRouterInterfaceService(client),
		SecurityGroup:       NewSecurityGroupService(client),
		SecurityGroupRule:   NewSecurityGroupRuleService(client),
	}
}
