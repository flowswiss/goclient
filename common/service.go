package common

import "github.com/flowswiss/goclient/v2/core"

type Service struct {
	Commitment *CommitmentService
	Location   *LocationService
	Module     *ModuleService
	Order      *OrderService
	Product    *ProductService
	Quota      *QuotaService
}

func NewService(client *core.Client) *Service {
	return &Service{
		Commitment: NewCommitmentService(client),
		Location:   NewLocationService(client),
		Module:     NewModuleService(client),
		Order:      NewOrderService(client),
		Product:    NewProductService(client),
		Quota:      NewQuotaService(client),
	}
}
