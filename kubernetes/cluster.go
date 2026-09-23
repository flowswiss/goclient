package kubernetes

import (
	"context"
	"encoding/json"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type ClusterService struct {
	client *core.Client
}

func NewClusterService(client *core.Client) *ClusterService {
	return &ClusterService{
		client: client,
	}
}

func (c ClusterService) List(ctx context.Context, cursor core.Cursor) (list common.List[Cluster], err error) {
	list.Pagination, err = c.client.List(ctx, getClusterPath(), cursor, &list.Items)
	return
}

type ClusterCreateReq struct {
	Name             string                 `json:"name"`
	LocationID       int                    `json:"location_id"`
	NetworkID        int                    `json:"network_id"`
	Worker           ClusterWorkerCreateReq `json:"worker"`
	AttachExternalIP bool                   `json:"attach_external_ip"`
}

type ClusterWorkerCreateReq struct {
	ProductID int `json:"product_id"`
	Count     int `json:"count"`
}

func (c ClusterService) Create(ctx context.Context, req ClusterCreateReq) (order common.Ordering, err error) {
	err = c.client.Create(ctx, getClusterPath(), req, &order)
	return
}

type ClusterGetReq struct {
	ID uint `json:"-"`
}

func (c ClusterService) Get(ctx context.Context, req ClusterGetReq) (cluster Cluster, err error) {
	err = c.client.Get(ctx, getSpecificClusterPath(req.ID), &cluster)
	return
}

type ClusterUpdateReq struct {
	ID uint `json:"-"`

	Name *string `json:"name,omitempty"`
}

func (c ClusterService) Update(ctx context.Context, req ClusterUpdateReq) (cluster Cluster, err error) {
	err = c.client.Update(ctx, getSpecificClusterPath(req.ID), req, &cluster)
	return
}

type ClusterDeleteReq struct {
	ID uint `json:"-"`
}

func (c ClusterService) Delete(ctx context.Context, req ClusterDeleteReq) (err error) {
	err = c.client.Delete(ctx, getSpecificClusterPath(req.ID))
	return
}

func (c ClusterService) GetKubeConfig(ctx context.Context, req ClusterGetReq) (
	kubeConfig ClusterKubeConfig,
	err error,
) {
	err = c.client.Get(ctx, getClusterKubeConfigPath(req.ID), &kubeConfig)
	return
}

func (c ClusterService) GetConfiguration(ctx context.Context, req ClusterGetReq) (
	config ClusterConfiguration,
	err error,
) {
	err = c.client.Get(ctx, getClusterConfigurationPath(req.ID), &config)
	return
}

type ClusterConfigurationReq struct {
	ID uint `json:"-"`

	VersionID int             `json:"version_id"`
	Variables json.RawMessage `json:"variables"`
}

func (c ClusterService) UpdateConfiguration(
	ctx context.Context,
	req ClusterConfigurationReq,
) (config ClusterConfiguration, err error) {
	err = c.client.Set(ctx, getClusterConfigurationPath(req.ID), req, &config)
	return
}

type ClusterUpdateFlavorReq struct {
	ID uint `json:"-"`

	Worker ClusterWorkerUpdateReq `json:"worker"`
}

type ClusterWorkerUpdateReq struct {
	ProductID int `json:"product_id"`
	Count     int `json:"count"`
}

func (c ClusterService) UpdateFlavor(ctx context.Context, req ClusterUpdateFlavorReq) (
	cluster Cluster,
	err error,
) {
	err = c.client.Update(ctx, getClusterFlavorPath(req.ID), req, &cluster)
	return
}

type ClusterPerformActionReq struct {
	ID uint `json:"-"`

	Action string `json:"action"`
}

func (c ClusterService) Perform(ctx context.Context, req ClusterPerformActionReq) (
	cluster Cluster,
	err error,
) {
	err = c.client.Create(ctx, getClusterActionPath(req.ID), req, &cluster)
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

func getSpecificClusterPath(id uint) string {
	return core.Join(clusterSegment, id)
}

func getClusterKubeConfigPath(id uint) string {
	return core.Join(clusterSegment, id, clusterKubeConfigSegment)
}

func getClusterConfigurationPath(id uint) string {
	return core.Join(clusterSegment, id, clusterConfigurationSegment)
}

func getClusterFlavorPath(id uint) string {
	return core.Join(clusterSegment, id, clusterFlavorSegment)
}

func getClusterActionPath(id uint) string {
	return core.Join(clusterSegment, id, clusterActionSegment)
}
