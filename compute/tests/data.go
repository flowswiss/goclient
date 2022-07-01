package computetests

import (
	commontests "github.com/flowswiss/goclient/common/tests"
	"github.com/flowswiss/goclient/compute"
)

const (
	LoadBalancerData       = `{"id": 1, "name": "lb-test", "location": ` + commontests.LocationData + `, "product": ` + commontests.ProductData + `, "status": ` + LoadBalancerStatusData + `, "networks": []}`
	LoadBalancerStatusData = `{"id": 1, "name": "Active", "key": "active"}`
)

var (
	LoadBalancer = compute.LoadBalancer{
		ID:       1,
		Name:     "lb-test",
		Location: commontests.Location,
		Product:  commontests.Product,
		Status:   LoadBalancerStatus,
		Networks: []compute.LoadBalancerNetworkAttachment{},
	}
	LoadBalancerStatus = compute.LoadBalancerStatus{
		ID:   1,
		Name: "Active",
		Key:  "active",
	}
	LoadBalancerPool = compute.LoadBalancerPool{
		ID:             0,
		Name:           "",
		Status:         LoadBalancerStatus,
		EntryProtocol:  compute.LoadBalancerProtocol{},
		TargetProtocol: compute.LoadBalancerProtocol{},
		EntryPort:      0,
		Algorithm:      compute.LoadBalancerAlgorithm{},
		StickySession:  false,
		HealthCheck:    compute.LoadBalancerHealthCheck{},
	}
	LoadBalancerMember = compute.LoadBalancerMember{
		ID:      0,
		Name:    "",
		Address: "",
		Port:    0,
		Status:  LoadBalancerStatus,
	}
)
