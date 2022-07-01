package commontests

import (
	"github.com/flowswiss/goclient/common"
)

const (
	BriefLocationData = `{"id": 1, "name": "ALP1"}`
	LocationData      = `{"id": 1, "name": "ALP1", "key": "key-alp1", "city": "Lucerne", "available_modules": [` + BriefModuleData + `]}`

	BriefModuleData = `{"id": 2, "name": "Compute"}`
	ModuleData      = `{"id": 2, "name": "Compute", "parent": ` + BriefModuleData + `, "sorting": 1, "locations": [` + BriefLocationData + `]}`

	OrderData    = `{"id": 1, "status": {"id": 1, "name": "created"}}`
	OrderingData = `{"ref": "https://api.flow.swiss/v4/orders/7"}`
	ProductData  = `{"id":13,"product_name":"macmini.2018.6-64-1024","type":{"id":5,"name":"Mac Bare Metal Device","key":"bare-metal-device"},"visibility":"public","usage_cycle":{"id":2,"name":"Hour","duration":1},"items":[{"id":12,"name":"Processor","description":"Core (3.6 GHz i7)","amount":6},{"id":13,"name":"Memory","description":"GB (2666 MHz DDR4)","amount":64},{"id":14,"name":"Storage","description":"GB (PCIe-based SSD)","amount":1024}],"children":[],"price":499,"availability":[{"location":` + BriefLocationData + `,"available":1}],"category":null,"deployment_fees":[{"location":` + BriefLocationData + `,"price":50,"free_deployments":1}]}`
)

var (
	BriefLocation = common.Location{
		ID:   1,
		Name: "ALP1",
	}
	Location = common.Location{
		ID:      1,
		Name:    "ALP1",
		Key:     "key-alp1",
		City:    "Lucerne",
		Modules: []common.Module{BriefModule},
	}

	BriefModule = common.Module{
		ID:   2,
		Name: "Compute",
	}
	Module = common.Module{
		ID:        2,
		Name:      "Compute",
		Parent:    &BriefModule,
		Sorting:   1,
		Locations: []common.Location{BriefLocation},
	}

	Order = common.Order{
		ID: 1,
		Status: common.OrderStatus{
			ID:   1,
			Name: "created",
		},
	}
	Ordering = common.Ordering{
		Ref: "https://api.flow.swiss/v4/orders/7",
	}

	Product = common.Product{
		ID:   13,
		Name: "macmini.2018.6-64-1024",
		Type: common.ProductType{
			ID:   5,
			Name: "Mac Bare Metal Device",
			Key:  "bare-metal-device",
		},
		Visibility: "public",
		UsageCycle: common.ProductUsageCycle{
			ID:       2,
			Name:     "Hour",
			Duration: 1,
		},
		Items: []common.ProductItem{
			{ID: 12, Name: "Processor", Description: "Core (3.6 GHz i7)", Amount: 6},
			{ID: 13, Name: "Memory", Description: "GB (2666 MHz DDR4)", Amount: 64},
			{ID: 14, Name: "Storage", Description: "GB (PCIe-based SSD)", Amount: 1024},
		},
		Price: 499,
		Availability: []common.ProductAvailability{
			{Location: BriefLocation, Available: 1},
		},
		Category: "",
		DeploymentFees: []common.DeploymentFee{
			{Location: BriefLocation, Price: 50, FreeDeployments: 1},
		},
	}
)
