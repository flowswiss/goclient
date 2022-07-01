package common

import (
	"context"

	"github.com/flowswiss/goclient"
)

type Location struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Key     string   `json:"key"`
	City    string   `json:"city"`
	Modules []Module `json:"available_modules"`
}

type LocationList struct {
	goclient.Pagination
	Items []Location
}

type LocationService struct {
	client goclient.Client
}

func NewLocationService(client goclient.Client) LocationService {
	return LocationService{client: client}
}

func (l LocationService) List(ctx context.Context, cursor goclient.Cursor) (list LocationList, err error) {
	list.Pagination, err = l.client.List(ctx, getLocationsPath(), cursor, &list.Items)
	return
}

func (l LocationService) Get(ctx context.Context, id int) (location Location, err error) {
	err = l.client.Get(ctx, getSpecificLocationPath(id), &location)
	return
}

const locationsSegment = "/v4/entities/locations"

func getLocationsPath() string {
	return locationsSegment
}

func getSpecificLocationPath(locationID int) string {
	return goclient.Join(locationsSegment, locationID)
}
