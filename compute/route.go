package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type Route struct {
	Id          int    `json:"id"`
	Destination string `json:"destination"`
	NextHop     string `json:"nexthop"`
}

type RouteList struct {
	Items      []Route
	Pagination goclient.Pagination
}

type RouteCreate struct {
	Name        string `json:"name"`
	Destination string `json:"destination"`
	NextHop     string `json:"nexthop"`
}

type RouteService struct {
	client   goclient.Client
	routerId int
}

func NewRouteService(client goclient.Client, routerId int) RouteService {
	return RouteService{client: client, routerId: routerId}
}

func (r RouteService) List(ctx context.Context, cursor goclient.Cursor) (list RouteList, err error) {
	list.Pagination, err = r.client.List(ctx, getRoutesPath(r.routerId), cursor, &list.Items)
	return
}

func (r RouteService) Create(ctx context.Context, body RouteCreate) (route Route, err error) {
	err = r.client.Create(ctx, getRoutesPath(r.routerId), body, &route)
	return
}

func (r RouteService) Delete(ctx context.Context, id int) (err error) {
	err = r.client.Delete(ctx, getSpecificRoutePath(r.routerId, id))
	return
}

const routesSegment = "routes"

func getRoutesPath(routerId int) string {
	return goclient.Join(routersSegment, routerId, routesSegment)
}

func getSpecificRoutePath(routerId, routeId int) string {
	return goclient.Join(routersSegment, routerId, routesSegment, routeId)
}
