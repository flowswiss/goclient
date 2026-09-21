package macbaremetal

import (
	"github.com/flowswiss/goclient/v2/common"
)

const (
	DirectionIngress = "ingress"
	DirectionEgress  = "egress"

	ProtocolAny  = -1
	ProtocolICMP = 1
	ProtocolTCP  = 6
	ProtocolUDP  = 17
)

// Devices

type Device struct {
	ID                int                        `json:"id"`
	Name              string                     `json:"name"`
	Location          common.Location            `json:"location"`
	Product           common.Product             `json:"product"`
	Status            DeviceStatus               `json:"status"`
	OperatingSystem   DeviceOperatingSystem      `json:"operating_system"`
	Network           Network                    `json:"network"`
	Hostname          string                     `json:"hostname"`
	NetworkInterfaces []AttachedNetworkInterface `json:"network_interfaces"`
	Price             float64                    `json:"price"`
	MetalControl      string                     `json:"metal_control"`
	MetalControlTools string                     `json:"metal_control_tools"`
}

type DeviceOperatingSystem struct {
	OS      string `json:"os"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AttachedNetworkInterface struct {
	ID        int    `json:"id"`
	PrivateIP string `json:"private_ip"`
	PublicIP  string `json:"public_ip"`
}

type DeviceVNCConnection struct {
	Ref string `json:"ref"`
}

type DeviceAction struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}

type DeviceStatus struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Key     string         `json:"key"`
	Actions []DeviceAction `json:"actions"`
}

type DeviceWorkflow struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}

// Elastic IPs

type ElasticIPAttachment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ElasticIP struct {
	ID         int                 `json:"id"`
	Product    common.BriefProduct `json:"product"`
	Location   common.Location     `json:"location"`
	Price      float64             `json:"price"`
	PublicIP   string              `json:"public_ip"`
	PrivateIP  string              `json:"private_ip"`
	Attachment ElasticIPAttachment `json:"attached_device"`
}

// Networks

type Network struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	Subnet              string          `json:"cidr"`
	Location            common.Location `json:"location"`
	DomainName          string          `json:"domain_name"`
	DomainNameServers   []string        `json:"domain_name_servers"`
	AllocationPoolStart string          `json:"allocation_pool_start"`
	AllocationPoolEnd   string          `json:"allocation_pool_end"`
	GatewayIP           string          `json:"gateway_ip"`
	UsedIPs             int             `json:"used_ips"`
	TotalIPs            int             `json:"total_ips"`
}

type NetworkInterface struct {
	ID                int           `json:"id"`
	PrivateIP         string        `json:"private_ip"`
	MacAddress        string        `json:"mac_address"`
	Network           Network       `json:"network"`
	SecurityGroup     SecurityGroup `json:"security_group"`
	AttachedElasticIP ElasticIP     `json:"attached_elastic_ip"`
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

type RouterInterface struct {
	ID        int     `json:"id"`
	PrivateIP string  `json:"private_ip"`
	Network   Network `json:"network"`
}

// Security Groups

type SecurityGroup struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Default     bool    `json:"default"`
	Network     Network `json:"network"`
}

type SecurityGroupRule struct {
	ID        int    `json:"id"`
	Direction string `json:"direction"`
	Protocol  int    `json:"protocol"`
	FromPort  int    `json:"from_port"`
	ToPort    int    `json:"to_port"`
	ICMPType  int    `json:"icmp_type"`
	ICMPCode  int    `json:"icmp_code"`
	IPRange   string `json:"ip_range"`
}
