package compute

import (
	"github.com/flowswiss/goclient/v2/common"
)

const (
	LoadBalancerStatusActive   = 1
	LoadBalancerStatusDisabled = 2
	LoadBalancerStatusWorking  = 3
	LoadBalancerStatusDegraded = 4
	LoadBalancerStatusError    = 5

	DirectionIngress = "ingress"
	DirectionEgress  = "egress"

	ProtocolAny  = -1
	ProtocolICMP = 1
	ProtocolTCP  = 6
	ProtocolUDP  = 17

	ServerStatusRunning   = 1
	ServerStatusStopped   = 2
	ServerStatusSuspended = 3
	ServerStatusStarting  = 4
	ServerStatusStopping  = 5
	ServerStatusError     = 6
	ServerStatusUpgrading = 9

	ApplicationStatusUnhealthy = 7
	ApplicationStatusHealthy   = 8

	ClusterStatusHealthy     = 14
	ClusterStatusWorking     = 15
	ClusterStatusUnhealthy   = 16
	ClusterStatusUnavailable = 17

	TaskStatusCreating   = 18
	TaskStatusNeutral    = 10
	TaskStatusCordoned   = 11
	TaskStatusDraining   = 12
	TaskStatusRebuilding = 13

	SnapshotStatusAvailable = 1
	SnapshotStatusCreating  = 2
	SnapshotStatusError     = 3

	VolumeStatusAvailable = 1
	VolumeStatusInUse     = 2
	VolumeStatusWorking   = 3
	VolumeStatusError     = 4
)

// Certificates

type Certificate struct {
	ID       int                `json:"id"`
	Name     string             `json:"name"`
	Location common.Location    `json:"location"`
	Type     string             `json:"type"`
	Details  CertificateDetails `json:"certificate"`
}

type CertificateDetails struct {
	Subject   map[string]string `json:"subject"`
	Issuer    map[string]string `json:"issuer"`
	ValidFrom common.Time       `json:"valid_from"`
	ValidTo   common.Time       `json:"valid_to"`
	Serial    string            `json:"serial"`
}

// Elastic IPs

type ElasticIPProduct = common.BriefProduct

type ElasticIPAttachment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ElasticIP struct {
	ID         int                 `json:"id"`
	Product    ElasticIPProduct    `json:"product"`
	Location   common.Location     `json:"location"`
	Price      float64             `json:"price"`
	PublicIP   string              `json:"public_ip"`
	PrivateIP  string              `json:"private_ip"`
	Attachment ElasticIPAttachment `json:"attached_instance"`
}

// Images

type Image struct {
	ID                 int              `json:"id"`
	OperatingSystem    string           `json:"os"`
	Version            string           `json:"version"`
	Key                string           `json:"key"`
	Category           string           `json:"category"`
	Type               string           `json:"type"`
	Username           string           `json:"username"`
	MinRootDiskSize    int              `json:"min_root_disk_size"`
	Sorting            int              `json:"sorting"`
	RequiredLicenses   []common.Product `json:"required_licenses"`
	AvailableLocations []int            `json:"available_locations"`
}

// Key Paris

type KeyPair struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
}

// Load Balancers

type LoadBalancer struct {
	ID       int                             `json:"id"`
	Name     string                          `json:"name"`
	Location common.Location                 `json:"location"`
	Product  common.Product                  `json:"product"`
	Status   LoadBalancerStatus              `json:"status"`
	Networks []LoadBalancerNetworkAttachment `json:"networks"`
}

type LoadBalancerNetworkAttachment struct {
	Network
	Interfaces []AttachedLoadBalancerInterface `json:"network_interfaces"`
}

type AttachedLoadBalancerInterface struct {
	ID        int    `json:"id"`
	PrivateIP string `json:"private_ip"`
	PublicIP  string `json:"public_ip"`
}

type LoadBalancerAlgorithm struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type LoadBalancerProtocol struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type LoadBalancerHealthCheckType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type LoadBalancerStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type LoadBalancerMember struct {
	ID      int                `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Port    int                `json:"port"`
	Status  LoadBalancerStatus `json:"status"`
}

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

type LoadBalancerHealthCheck struct {
	Type               LoadBalancerHealthCheckType `json:"type"`
	HTTPMethod         string                      `json:"http_method"`
	HTTPPath           string                      `json:"http_path"`
	Interval           int                         `json:"interval"`
	Timeout            int                         `json:"timeout"`
	HealthyThreshold   int                         `json:"healthy_threshold"`
	UnhealthyThreshold int                         `json:"unhealthy_threshold"`
}

// Networks

type Network struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	CIDR                string          `json:"cidr"`
	Location            common.Location `json:"location"`
	DomainNameServers   []string        `json:"domain_name_servers"`
	AllocationPoolStart string          `json:"allocation_pool_start"`
	AllocationPoolEnd   string          `json:"allocation_pool_end"`
	GatewayIP           string          `json:"gateway_ip"`
	UsedIPs             int             `json:"used_ips"`
	TotalIPs            int             `json:"total_ips"`
}

type NetworkBrief struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	CIDR     string          `json:"cidr"`
	Location common.Location `json:"location"`
	UsedIPs  int             `json:"used_ips"`
	TotalIPs int             `json:"total_ips"`
}

type NetworkInterface struct {
	ID                int             `json:"id"`
	PrivateIP         string          `json:"private_ip"`
	MacAddress        string          `json:"mac_address"`
	Network           Network         `json:"network"`
	AttachedElasticIP ElasticIP       `json:"attached_elastic_ip"`
	SecurityGroups    []SecurityGroup `json:"security_groups"`
	Security          bool            `json:"security"`
}

// Routes

type Route struct {
	ID          int    `json:"id"`
	Destination string `json:"destination"`
	NextHop     string `json:"nexthop"`
}

// Router

type Router struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Location    common.Location `json:"location"`
	Public      bool            `json:"public"`
	SourceNAT   bool            `json:"snat"`
	PublicIP    string          `json:"public_ip"`
}

// Router Interfaces

type RouterInterface struct {
	ID        int     `json:"id"`
	PrivateIP string  `json:"private_ip"`
	Network   Network `json:"network"`
}

// Security Groups

type SecurityGroup struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Location    common.Location `json:"location"`
	Default     bool            `json:"default"`
	Immutable   bool            `json:"immutable"`
}

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

// Servers

type Server struct {
	ID       int                       `json:"id"`
	Name     string                    `json:"name"`
	Status   ServerStatus              `json:"status"`
	Image    Image                     `json:"image"`
	Product  common.Product            `json:"product"`
	Location common.Location           `json:"location"`
	Networks []ServerNetworkAttachment `json:"networks"`
	KeyPair  KeyPair                   `json:"key_pair"`
}

type ServerNetworkAttachment struct {
	Network
	Interfaces []AttachedNetworkInterface `json:"network_interfaces"`
}

type AttachedNetworkInterface struct {
	ID        int    `json:"id"`
	PrivateIP string `json:"private_ip"`
	PublicIP  string `json:"public_ip"`
}

type ServerAction struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}

type ServerStatus struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Key     string         `json:"key"`
	Actions []ServerAction `json:"actions"`
}

// Snapshots

type Snapshot struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Size      int            `json:"size"`
	Status    SnapshotStatus `json:"status"`
	Volume    Volume         `json:"volume"`
	Product   common.Product `json:"product"`
	CreatedAt common.Time    `json:"created_at"`
}

type SnapshotStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Volumes

type VolumeStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Volume struct {
	ID           int             `json:"id"`
	Product      common.Product  `json:"product"`
	Location     common.Location `json:"location"`
	Status       VolumeStatus    `json:"status"`
	Name         string          `json:"name"`
	Size         int             `json:"size"`
	SerialNumber string          `json:"serial"`
	Snapshots    int             `json:"snapshots"`
	Bootable     bool            `json:"bootable"`
	RootVolume   bool            `json:"root_volume"`
	AttachedTo   Server          `json:"instance"`
	CreatedAt    common.Time     `json:"created_at"`
}

// Connections

type Connection struct {
	ID             int                      `json:"id"`
	Type           string                   `json:"type"`
	Name           string                   `json:"name"`
	MTU            int                      `json:"mtu"`
	Location       common.Location          `json:"location"`
	LocalEndpoint  ConnectionLocalEndpoint  `json:"local_endpoint"`
	RemoteEndpoint ConnectionRemoteEndpoint `json:"remote_endpoint"`
	Status         ConnectionStatus         `json:"status"`
	PSK            int                      `json:"psk"`
	IKEPolicy      struct {
		Lifetime                int                   `json:"lifetime"`
		AuthenticationAlgorithm ConnectionNamedEntity `json:"authentication_algorithm"`
		EncryptionAlgorithm     ConnectionNamedEntity `json:"encryption_algorithm"`
		DiffieHellmanGroup      ConnectionNamedEntity `json:"diffie_hellman_group"`
		IkeVersion              ConnectionNamedEntity `json:"ike_version"`
	} `json:"ike_policy"`
	IPSecPolicy struct {
		Lifetime                int                   `json:"lifetime"`
		AuthenticationAlgorithm ConnectionNamedEntity `json:"authentication_algorithm"`
		EncryptionAlgorithm     ConnectionNamedEntity `json:"encryption_algorithm"`
		DiffieHellmanGroup      ConnectionNamedEntity `json:"diffie_hellman_group"`
		EncapsulationMode       ConnectionNamedEntity `json:"encapsulation_mode"`
		TransformProtocol       ConnectionNamedEntity `json:"transform_protocol"`
	} `json:"ipsec_policy"`
	DeadPeerDetection struct {
		Action   ConnectionNamedEntity `json:"action"`
		Interval int                   `json:"interval"`
		Timeout  int                   `json:"timeout"`
	} `json:"dead_peer_detection"`
	Initiator ConnectionNamedEntity `json:"initiator"`
	Product   common.Product        `json:"product"`
}

type ConnectionEndpoint struct {
	PrivateNetworkID int `json:"private_network_id"`
}

type ConnectionExternalEndpoint struct {
	PeerIP string   `json:"peer_ip"`
	CIDRs  []string `json:"cidrs"`
}

type ConnectionIPSecPolicy struct {
	Lifetime                  int `json:"lifetime"`
	AuthenticationAlgorithmId int `json:"authentication_algorithm_id"`
	EncryptionAlgorithmId     int `json:"encryption_algorithm_id"`
	DiffieHellmanGroupId      int `json:"diffie_hellman_group_id"`
	EncapsulationModeId       int `json:"encapsulation_mode_id"`
	TransformProtocolId       int `json:"transform_protocol_id"`
}

type ConnectionDeadPeerDetection struct {
	ActionID int `json:"action_id"`
	Interval int `json:"interval"`
	Timeout  int `json:"timeout"`
}

type ConnectionIKEPolicy struct {
	Lifetime                  int `json:"lifetime"`
	AuthenticationAlgorithmId int `json:"authentication_algorithm_id"`
	EncryptionAlgorithmId     int `json:"encryption_algorithm_id"`
	DiffieHellmanGroupId      int `json:"diffie_hellman_group_id"`
	IkeVersionId              int `json:"ike_version_id"`
}

type ConnectionLocalEndpoint struct {
	Network NetworkBrief `json:"network"`
	Router  Router       `json:"router"`
}

type ConnectionRemoteEndpoint struct {
	PeerIP string   `json:"peer_ip"`
	CIDRs  []string `json:"cidrs"`
}

type ConnectionStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ConnectionNamedEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
