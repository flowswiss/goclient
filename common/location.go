package common

import (
	"context"

	"github.com/flowswiss/goclient/v2/core"
)

type Location struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Key     string   `json:"key"`
	City    string   `json:"city"`
	Modules []Module `json:"available_modules"`
}

type LocationService struct {
	client *core.Client
}

func NewLocationService(client *core.Client) *LocationService {
	return &LocationService{client: client}
}

func (l LocationService) List(ctx context.Context, cursor core.Cursor) (list List[Location], err error) {
	list.Pagination, err = l.client.List(ctx, getLocationsPath(), cursor, &list.Items)
	return
}

type LocationGetReq struct {
	ID uint `json:"-"`
}

func (l LocationService) Get(ctx context.Context, req LocationGetReq) (location Location, err error) {
	err = l.client.Get(ctx, getSpecificLocationPath(req.ID), &location)
	return
}

const locationsSegment = "/v4/entities/locations"

func getLocationsPath() string {
	return locationsSegment
}

func getSpecificLocationPath(locationID uint) string {
	return core.Join(locationsSegment, locationID)
}
