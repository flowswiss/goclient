package kubernetes

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type NodeService struct {
	client *core.Client
}

func NewNodeService(client *core.Client) *NodeService {
	return &NodeService{
		client: client,
	}
}

type NodeListReq struct {
	ClusterID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (n NodeService) List(ctx context.Context, req NodeListReq) (list common.List[Node], err error) {
	list.Pagination, err = n.client.List(ctx, getNodesPath(req.ClusterID), req.Cursor, &list.Items)
	return
}

type NodeDeleteReq struct {
	ClusterID uint `json:"-"`
	NodeID    uint `json:"-"`
}

func (n NodeService) Delete(ctx context.Context, req NodeDeleteReq) (err error) {
	err = n.client.Delete(ctx, getSpecificNodePath(req.ClusterID, req.NodeID))
	return
}

type NodePerformReq struct {
	ClusterID uint `json:"-"`
	NodeID    uint `json:"-"`

	Action string `json:"action"`
}

func (n NodeService) Perform(ctx context.Context, req NodePerformReq) (node Node, err error) {
	err = n.client.Create(ctx, getNodeActionPath(req.ClusterID, req.NodeID), req, &node)
	return
}

const (
	nodesSegment      = "nodes"
	nodeActionSegment = "action"
)

func getNodesPath(clusterID uint) string {
	return core.Join(clustersSegment, clusterID, nodesSegment)
}

func getSpecificNodePath(clusterID, nodeID uint) string {
	return core.Join(clustersSegment, clusterID, nodesSegment, nodeID)
}

func getNodeActionPath(clusterID, nodeID uint) string {
	return core.Join(clustersSegment, clusterID, nodesSegment, nodeID, nodeActionSegment)
}
