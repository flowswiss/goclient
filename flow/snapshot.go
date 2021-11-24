package flow

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type SnapshotService interface {
	Get(ctx context.Context, snapshotId Id) (*Snapshot, *Response, error)
	List(ctx context.Context, options PaginationOptions) ([]*Snapshot, *Response, error)
	Create(ctx context.Context, data *SnapshotCreate) (*Snapshot, *Response, error)
	Delete(ctx context.Context, snapshotId Id) (*Response, error)
}

type SnaphostStatusKey = string

const (
	SnapshotStatusKeyAvailable SnaphostStatusKey = "available"
	SnapshotStatusKeyCreating  SnaphostStatusKey = "creating"
	SnapshotStatusKeyError     SnaphostStatusKey = "error"
)

type Status struct {
	Id   Id                `json:"id"`
	Name string            `json:"name"`
	Key  string `json:"key"`
}

type Snapshot struct {
	Id        Id       `json:"id"`
	Name      string   `json:"name"`
	Size      int      `json:"size"`
	Status    Status   `json:"status"`
	Volume    Volume   `json:"volume"`
	Product   Product  `json:"product"`
	CreatedAt DateTime `json:"created_at"`
}

func (snapshot *Snapshot) IsAvailable() (bool, error) {
	switch snapshot.Status.Key {
	case SnapshotStatusKeyAvailable:
		return true, nil
	case SnapshotStatusKeyCreating:
		return false, nil
	case SnapshotStatusKeyError:
		return false, errors.New(fmt.Sprintf("Snapshot with id %d has errored state", snapshot.Id))
	}
	return false, errors.New(fmt.Sprintf("Snapshot with id %d has unknown status '%s'", snapshot.Id, snapshot.Status.Key))
}

type SnapshotCreate struct {
	Name     string `json:"name"`
	VolumeId Id     `json:"volume_id"`
}

type snapshotService struct {
	client *Client
}

func (s *snapshotService) Get(ctx context.Context, snapshotId Id) (*Snapshot, *Response, error) {
	p := fmt.Sprintf("/v4/compute/snapshots/%d", snapshotId)

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	var val *Snapshot

	res, err := s.client.Do(req, &val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
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
