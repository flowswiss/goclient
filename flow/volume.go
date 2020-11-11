package flow

import (
	"context"
	"fmt"
	"net/http"
)

const (
	VolumeStatusAvailable = Id(1)
	VolumeStatusInUse     = Id(2)
	VolumeStatusWorking   = Id(3)
	VolumeStatusError     = Id(4)
)

type VolumeService interface {
	List(ctx context.Context, options PaginationOptions) ([]*Volume, *Response, error)
	Get(ctx context.Context, id Id) (*Volume, *Response, error)
	Create(ctx context.Context, data *VolumeCreate) (*Volume, *Response, error)
	Delete(ctx context.Context, id Id) (*Response, error)

	Expand(ctx context.Context, id Id, data *VolumeExpand) (*Volume, *Response, error)
}

type VolumeStatus struct {
	Id   Id     `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Volume struct {
	Id           Id           `json:"id"`
	Product      Product      `json:"product"`
	Location     Location     `json:"location"`
	Status       VolumeStatus `json:"status"`
	Name         string       `json:"name"`
	Size         int          `json:"size"`
	SerialNumber string       `json:"serial"`
	Snapshots    int          `json:"snapshots"`
	Bootable     bool         `json:"bootable"`
	RootVolume   bool         `json:"root_volume"`
	AttachedTo   *Server      `json:"instance"`
	CreatedAt    DateTime     `json:"created_at"`
}

type VolumeCreate struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	LocationId Id     `json:"location_id"`
	SnapshotId Id     `json:"snapshot_id,omitempty"`
	InstanceId Id     `json:"instance_id,omitempty"`
}

type VolumeExpand struct {
	Size int `json:"size"`
}

type volumeService struct {
	client *Client
}

func (s *volumeService) List(ctx context.Context, options PaginationOptions) ([]*Volume, *Response, error) {
	p, err := addOptions("/v4/compute/volumes", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	var val []*Volume

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *volumeService) Get(ctx context.Context, id Id) (*Volume, *Response, error) {
	p := fmt.Sprintf("/v4/compute/volumes/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	val := &Volume{}

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *volumeService) Create(ctx context.Context, data *VolumeCreate) (*Volume, *Response, error) {
	p := "/v4/compute/volumes"

	req, err := s.client.NewRequest(ctx, http.MethodPost, p, data)
	if err != nil {
		return nil, nil, err
	}

	val := &Volume{}

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *volumeService) Delete(ctx context.Context, id Id) (*Response, error) {
	p := fmt.Sprintf("/v4/compute/volumes/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, p, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

func (s *volumeService) Expand(ctx context.Context, id Id, expand *VolumeExpand) (*Volume, *Response, error) {
	p := fmt.Sprintf("/v4/compute/volumes/%d/upgrade", id)

	req, err := s.client.NewRequest(ctx, http.MethodPost, p, expand)
	if err != nil {
		return nil, nil, err
	}

	val := &Volume{}

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}
