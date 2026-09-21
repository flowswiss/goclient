package compute

import "github.com/flowswiss/goclient/v2/core"

type Service struct {
	Certificate         *CertificateService
	Connection          *ConnectionService
	ElasticIP           *ElasticIPService
	Image               *ImageService
	KeyPair             *KeyPairService
	LoadBalancer        *LoadBalancerService
	LoadBalancerEntity  *LoadBalancerEntityService
	LoadBalancerMember  *LoadBalancerMemberService
	LoadBalancerPool    *LoadBalancerPoolService
	Network             *NetworkService
	NetworkInterface    *NetworkInterfaceService
	Route               *RouteService
	Router              *RouterService
	RouterInterface     *RouterInterfaceService
	SecurityGroup       *SecurityGroupService
	SecurityGroupRule   *SecurityGroupRuleService
	Server              *ServerService
	ElasticIPAttachment *ElasticIPAttachmentService
	Snapshot            *SnapshotService
	Volume              *VolumeService
}

func New(client *core.Client) *Service {
	return &Service{
		Certificate:         NewCertificateService(client),
		Connection:          NewConnectionService(client),
		ElasticIP:           NewElasticIPService(client),
		ElasticIPAttachment: NewElasticIPAttachmentService(client),
		Image:               NewImageService(client),
		KeyPair:             NewKeyPairService(client),
		LoadBalancer:        NewLoadBalancerService(client),
		LoadBalancerEntity:  NewLoadBalancerEntityService(client),
		LoadBalancerMember:  NewLoadBalancerMemberService(client),
		LoadBalancerPool:    NewLoadBalancerPoolService(client),
		Network:             NewNetworkService(client),
		NetworkInterface:    NewNetworkInterfaceService(client),
		Route:               NewRouteService(client),
		Router:              NewRouterService(client),
		RouterInterface:     NewRouterInterfaceService(client),
		SecurityGroup:       NewSecurityGroupService(client),
		SecurityGroupRule:   NewSecurityGroupRuleService(client),
		Server:              NewServerService(client),
		Snapshot:            NewSnapshotService(client),
		Volume:              NewVolumeService(client),
	}
}
