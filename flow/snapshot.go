package flow

import (
	"context"
	"fmt"
	"net/http"
)

type SnapshotService interface {
	List(ctx context.Context, options PaginationOptions) ([]*Snapshot, *Response, error)
	Create(ctx context.Context, data *SnapshotCreate) (*Snapshot, *Response, error)
	Delete(ctx context.Context, snapshotId Id) (*Response, error)
}

type Snapshot struct {
	Id        Id       `json:"id"`
	Name      string   `json:"name"`
	Size      int      `json:"size"`
	Volume    Volume   `json:"volume"`
	Product   Product  `json:"product"`
	CreatedAt DateTime `json:"created_at"`
}

type SnapshotCreate struct {
	Name     string `json:"name"`
	VolumeId Id     `json:"volume_id"`
}

type snapshotService struct {
	client *Client
}

func (s *snapshotService) List(ctx context.Context, options PaginationOptions) ([]*Snapshot, *Response, error) {
	p, err := addOptions("/v4/compute/snapshots", options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	var val []*Snapshot

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *snapshotService) Create(ctx context.Context, data *SnapshotCreate) (*Snapshot, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "/v4/compute/snapshots", data)
	if err != nil {
		return nil, nil, err
	}

	val := &Snapshot{}

	res, err := s.client.Do(req, val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *snapshotService) Delete(ctx context.Context, snapshotId Id) (*Response, error) {
	p := fmt.Sprintf("/v4/compute/snapshots/%d", snapshotId)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, p, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
