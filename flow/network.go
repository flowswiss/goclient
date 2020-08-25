package flow

import (
	"context"
	"fmt"
	"net/http"
)

type NetworkService interface {
	List(ctx context.Context, options PaginationOptions) ([]*Network, *Response, error)
	Get(ctx context.Context, id Id) (*Network, *Response, error)
	Create(ctx context.Context, data *NetworkCreate) (*Network, *Response, error)
	Delete(ctx context.Context, id Id) (*Response, error)
}

type Network struct {
	Id                  Id       `json:"id"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Cidr                string   `json:"cidr"`
	Location            Location `json:"location"`
	DomainNameServers   []string `json:"domain_name_servers"`
	AllocationPoolStart string   `json:"allocation_pool_start"`
	AllocationPoolEnd   string   `json:"allocation_pool_end"`
	GatewayIp           string   `json:"gateway_ip"`
	UsedIps             int      `json:"used_ips"`
	TotalIps            int      `json:"total_ips"`
}

type NetworkCreate struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	LocationId          Id       `json:"location_id"`
	DomainNameServer    []string `json:"domain_name_server"`
	Cidr                string   `json:"cidr"`
	AllocationPoolStart string   `json:"allocation_pool_start"`
	AllocationPoolEnd   string   `json:"allocation_pool_end"`
	GatewayIp           string   `json:"gateway_ip"`
}

type networkService struct {
	client *Client
}

func (s *networkService) List(ctx context.Context, options PaginationOptions) ([]*Network, *Response, error) {
	p := "/v4/compute/networks"
	p, err := addOptions(p, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	var val []*Network

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *networkService) Get(ctx context.Context, id Id) (*Network, *Response, error) {
	p := fmt.Sprintf("/v4/compute/networks/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	val := &Network{}

	res, err := s.client.Do(req, val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *networkService) Create(ctx context.Context, data *NetworkCreate) (*Network, *Response, error) {
	p := "/v4/compute/networks"

	req, err := s.client.NewRequest(ctx, http.MethodPost, p, data)
	if err != nil {
		return nil, nil, err
	}

	val := &Network{}

	res, err := s.client.Do(req, val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *networkService) Delete(ctx context.Context, id Id) (*Response, error) {
	p := fmt.Sprintf("/v4/compute/networks/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, p, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
