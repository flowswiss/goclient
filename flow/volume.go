package flow

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type VolumeService interface {
	List(ctx context.Context, options PaginationOptions) ([]*Volume, *Response, error)
	Get(ctx context.Context, id Id) (*Volume, *Response, error)
	Create(ctx context.Context, data *VolumeCreate) (*Volume, *Response, error)
	Delete(ctx context.Context, id Id) (*Response, error)
}

type VolumeStatus struct {
	Id   Id     `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Volume struct {
	Id         Id           `json:"id"`
	Product    Product      `json:"product"`
	Location   Location     `json:"location"`
	Status     VolumeStatus `json:"status"`
	Name       string       `json:"name"`
	Size       int          `json:"size"`
	Snapshots  int          `json:"snapshots"`
	Bootable   bool         `json:"bootable"`
	RootVolume bool         `json:"root_volume"`
	AttachedTo Server       `json:"instance"`
	CreatedAt  time.Time    `json:"created_at"`
}

type VolumeCreate struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	LocationId Id     `json:"location_id"`
	SnapshotId int    `json:"snapshot_id,omitempty"`
	InstanceId int    `json:"instance_id,omitempty"`
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
