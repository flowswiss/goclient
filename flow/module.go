package flow

import (
	"context"
	"fmt"
	"net/http"
)

type ModuleService interface {
	List(ctx context.Context, options PaginationOptions) ([]*Module, *Response, error)
	Get(ctx context.Context, id Id) (*Module, *Response, error)
}

type Module struct {
	Id        Id          `json:"id"`
	Name      string      `json:"name"`
	Parent    *Module     `json:"parent"`
	Sorting   int         `json:"sorting"`
	Locations []*Location `json:"locations"`
}

func (m *Module) AvailableAt(location *Location) bool {
	for _, available := range m.Locations {
		if available.Id == location.Id {
			return true
		}
	}
	return false
}

type moduleService struct {
	client *Client
}

func (s *moduleService) List(ctx context.Context, options PaginationOptions) ([]*Module, *Response, error) {
	p, err := addOptions("/v4/entities/modules", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	var val []*Module

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *moduleService) Get(ctx context.Context, id Id) (*Module, *Response, error) {
	p := fmt.Sprintf("/v4/entities/modules/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	val := &Module{}

	res, err := s.client.Do(req, val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}
