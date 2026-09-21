package kubernetes

import (
	"encoding/json"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/compute"
)

// Clusters

type Cluster struct {
	ID            int                   `json:"id"`
	Name          string                `json:"name"`
	Location      common.Location       `json:"location"`
	Product       common.Product        `json:"product"`
	Network       compute.Network       `json:"network"`
	SecurityGroup compute.SecurityGroup `json:"security_group"`

	PublicAddress string `json:"public_address"`
	DNSName       string `json:"dns_name"`

	NodeCount struct {
		Current struct {
			ControlPlane int `json:"control-plane"`
			Worker       int `json:"worker"`
		} `json:"current"`

		Expected struct {
			ControlPlane int `json:"control-plane"`
			Worker       int `json:"worker"`
		} `json:"expected"`
	} `json:"node_count"`

	ExpectedPreset struct {
		ControlPlane common.Product `json:"control_plane"`
		Worker       common.Product `json:"worker"`
	} `json:"expected_preset"`

	Version ClusterVersion `json:"kubernetes_version"`

	Status ClusterStatus `json:"status"`
	Locked bool          `json:"locked"`

	KubeConfig struct {
		UpdatedAt common.Time `json:"updated_at"`
		ExpiresAt common.Time `json:"expires_at"`
	} `json:"kube_config"`
}

type ClusterVersion struct {
	ID           int              `json:"id"`
	Name         string           `json:"name"`
	Major        int              `json:"major"`
	Minor        int              `json:"minor"`
	Schema       json.RawMessage  `json:"schema"`
	UpgradePaths []ClusterVersion `json:"upgrade_paths"`
	HostImage    compute.Image    `json:"host_image"`
}

type ClusterStatus struct {
	ID      int             `json:"id"`
	Key     string          `json:"key"`
	Name    string          `json:"name"`
	Actions []ClusterAction `json:"actions"`
}

type ClusterAction struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}

type ClusterKubeConfig struct {
	KubeConfig string `json:"kube_config"`
}

type ClusterConfiguration struct {
	VersionID int             `json:"version_id"`
	Variables json.RawMessage `json:"variables"`
}

// Nodes

type Node struct {
	ID      int                             `json:"id"`
	Name    string                          `json:"name"`
	Roles   []NodeRole                      `json:"roles"`
	Product common.Product                  `json:"product"`
	Network compute.ServerNetworkAttachment `json:"network"`
	Status  NodeStatus                      `json:"status"`
}

type NodeRole struct {
	ID   int    `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type NodeStatus struct {
	ID      int          `json:"id"`
	Key     string       `json:"key"`
	Name    string       `json:"name"`
	Actions []NodeAction `json:"actions"`
}

type NodeAction struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}
