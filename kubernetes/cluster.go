package kubernetes

import (
	"context"
	"encoding/json"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
	"github.com/flowswiss/goclient/compute"
)

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

type ClusterList struct {
	Items      []Cluster
	Pagination goclient.Pagination
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

type ClusterCreate struct {
	Name             string              `json:"name"`
	LocationID       int                 `json:"location_id"`
	NetworkID        int                 `json:"network_id"`
	Worker           ClusterWorkerCreate `json:"worker"`
	AttachExternalIP bool                `json:"attach_external_ip"`
}

type ClusterWorkerCreate struct {
	ProductID int `json:"product_id"`
	Count     int `json:"count"`
}

type ClusterUpdate struct {
	Name string `json:"name,omitempty"`
}

type ClusterUpdateFlavor struct {
	Worker ClusterWorkerUpdate `json:"worker"`
}

type ClusterWorkerUpdate struct {
	ProductID int `json:"product_id"`
	Count     int `json:"count"`
}

type ClusterPerformAction struct {
	Action string `json:"action"`
}

type ClusterService struct {
	client goclient.Client
}

func NewClusterService(client goclient.Client) ClusterService {
	return ClusterService{
		client: client,
	}
}

func (c ClusterService) Nodes(clusterID int) NodeService {
	return NewNodeService(c.client, clusterID)
}

func (c ClusterService) Volumes(clusterID int) VolumeService {
	return NewVolumeService(c.client, clusterID)
}

func (c ClusterService) LoadBalancers(clusterID int) LoadBalancerService {
	return NewLoadBalancerService(c.client, clusterID)
}

func (c ClusterService) List(ctx context.Context, cursor goclient.Cursor) (list ClusterList, err error) {
	list.Pagination, err = c.client.List(ctx, getClusterPath(), cursor, &list.Items)
	return
}

func (c ClusterService) Create(ctx context.Context, body ClusterCreate) (order common.Ordering, err error) {
	err = c.client.Create(ctx, getClusterPath(), body, &order)
	return
}

func (c ClusterService) Get(ctx context.Context, id int) (cluster Cluster, err error) {
	err = c.client.Get(ctx, getSpecificClusterPath(id), &cluster)
	return
}

func (c ClusterService) Update(ctx context.Context, id int, body ClusterUpdate) (cluster Cluster, err error) {
	err = c.client.Update(ctx, getSpecificClusterPath(id), body, &cluster)
	return
}

func (c ClusterService) Delete(ctx context.Context, id int) (err error) {
	err = c.client.Delete(ctx, getSpecificClusterPath(id))
	return
}

func (c ClusterService) GetKubeConfig(ctx context.Context, id int) (kubeConfig ClusterKubeConfig, err error) {
	err = c.client.Get(ctx, getClusterKubeConfigPath(id), &kubeConfig)
	return
}

func (c ClusterService) GetConfiguration(ctx context.Context, id int) (config ClusterConfiguration, err error) {
	err = c.client.Get(ctx, getClusterConfigurationPath(id), &config)
	return
}

func (c ClusterService) UpdateConfiguration(ctx context.Context, id int, body ClusterConfiguration) (config ClusterConfiguration, err error) {
	err = c.client.Set(ctx, getClusterConfigurationPath(id), body, &config)
	return
}

func (c ClusterService) UpdateFlavor(ctx context.Context, id int, body ClusterUpdateFlavor) (cluster Cluster, err error) {
	err = c.client.Update(ctx, getClusterFlavorPath(id), body, &cluster)
	return
}

func (c ClusterService) PerformAction(ctx context.Context, id int, body ClusterPerformAction) (cluster Cluster, err error) {
	err = c.client.Create(ctx, getClusterActionPath(id), body, &cluster)
	return
}

const (
	clusterSegment              = "/v4/kubernetes/clusters"
	clusterKubeConfigSegment    = "kube-config"
	clusterConfigurationSegment = "configuration"
	clusterFlavorSegment        = "flavor"
	clusterActionSegment        = "action"
)

func getClusterPath() string {
	return clusterSegment
}

func getSpecificClusterPath(id int) string {
	return goclient.Join(clusterSegment, id)
}

func getClusterKubeConfigPath(id int) string {
	return goclient.Join(clusterSegment, id, clusterKubeConfigSegment)
}

func getClusterConfigurationPath(id int) string {
	return goclient.Join(clusterSegment, id, clusterConfigurationSegment)
}

func getClusterFlavorPath(id int) string {
	return goclient.Join(clusterSegment, id, clusterFlavorSegment)
}

func getClusterActionPath(id int) string {
	return goclient.Join(clusterSegment, id, clusterActionSegment)
}
