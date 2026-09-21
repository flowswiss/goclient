package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type RouteService struct {
	client *core.Client
}

func NewRouteService(client *core.Client) *RouteService {
	return &RouteService{client: client}
}

type RouteListReq struct {
	RouterID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (r RouteService) List(ctx context.Context, req RouteListReq) (list common.List[Route], err error) {
	list.Pagination, err = r.client.List(ctx, getRoutesPath(req.RouterID), req.Cursor, &list.Items)
	return
}

type RouteCreateReq struct {
	RouterID uint `json:"-"`

	Destination string `json:"destination"`
	NextHop     string `json:"nexthop"`
}

func (r RouteService) Create(ctx context.Context, req RouteCreateReq) (route Route, err error) {
	err = r.client.Create(ctx, getRoutesPath(req.RouterID), req, &route)
	return
}

type RouteDeleteReq struct {
	RouterID uint `json:"-"`
	RouteID  uint `json:"-"`
}

func (r RouteService) Delete(ctx context.Context, req RouteDeleteReq) (err error) {
	err = r.client.Delete(ctx, getSpecificRoutePath(req.RouterID, req.RouteID))
	return
}

const routesSegment = "routes"

func getRoutesPath(routerID uint) string {
	return core.Join(routersSegment, routerID, routesSegment)
}

func getSpecificRoutePath(routerID, routeID uint) string {
	return core.Join(routersSegment, routerID, routesSegment, routeID)
}
