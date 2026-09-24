package kubernetes

import "github.com/flowswiss/goclient/v2/core"

type Service struct {
	Cluster      *ClusterService
	LoadBalancer *LoadBalancerService
	Node         *NodeService
	Volume       *VolumeService
	Snapshot     *SnapshotService
}

func NewService(client *core.Client) *Service {
	return &Service{
		Cluster:      NewClusterService(client),
		LoadBalancer: NewLoadBalancerService(client),
		Node:         NewNodeService(client),
		Volume:       NewVolumeService(client),
		Snapshot:     NewSnapshotService(client),
	}
}
