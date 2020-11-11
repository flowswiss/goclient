package flow

import (
	"context"
	"fmt"
	"net/http"
)

type VolumeActionService interface {
	Attach(ctx context.Context, volumeId Id, data *VolumeAttach) (*Volume, *Response, error)
	Detach(ctx context.Context, volumeId Id, serverId Id) (*Response, error)
}

type VolumeAttach struct {
	ServerId Id `json:"instance_id"`
}

type volumeActionService struct {
	client *Client
}

func (s *volumeActionService) Attach(ctx context.Context, volumeId Id, data *VolumeAttach) (*Volume, *Response, error) {
	p := fmt.Sprintf("/v4/compute/volumes/%d/instances", volumeId)

	req, err := s.client.NewRequest(ctx, http.MethodPost, p, data)
	if err != nil {
		return nil, nil, err
	}

	val := &Volume{}

	res, err := s.client.Do(req, val)
	if err != nil {
		return nil, nil, err
	}

	return val, res, nil
}

func (s *volumeActionService) Detach(ctx context.Context, volumeId Id, serverId Id) (*Response, error) {
	p := fmt.Sprintf("/v4/compute/volumes/%d/instances/%d", volumeId, serverId)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, p, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
