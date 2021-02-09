package compute

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Router struct {
	Id          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Location    common.Location `json:"location"`
	Public      bool            `json:"public"`
	Snat        bool            `json:"snat"`
	PublicIp    string          `json:"public_ip"`
}

type RouterList struct {
	Items      []Router
	Pagination goclient.Pagination
}

type RouterCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LocationId  int    `json:"location_id"`
	Public      bool   `json:"public"`
}

type RouterUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

type RouterService struct {
	client goclient.Client
}

func NewRouterService(client goclient.Client) RouterService {
	return RouterService{client: client}
}

func (r RouterService) RouterInterfaces(routerId int) RouterInterfaceService {
	return NewRouterInterfaceService(r.client, routerId)
}

func (r RouterService) Routes(routerId int) RouteService {
	return NewRouteService(r.client, routerId)
}

func (r RouterService) List(ctx context.Context, cursor goclient.Cursor) (list RouterList, err error) {
	list.Pagination, err = r.client.List(ctx, getRoutersPath(), cursor, &list.Items)
	return
}

func (r RouterService) Get(ctx context.Context, id int) (router Router, err error) {
	err = r.client.Get(ctx, getSpecificRouterPath(id), &router)
	return
}

func (r RouterService) Create(ctx context.Context, body RouterCreate) (router Router, err error) {
	err = r.client.Create(ctx, getRoutersPath(), body, &router)
	return
}

func (r RouterService) Update(ctx context.Context, id int, body RouterUpdate) (router Router, err error) {
	err = r.client.Update(ctx, getSpecificRouterPath(id), body, &router)
	return
}

func (r RouterService) Delete(ctx context.Context, id int) (err error) {
	err = r.client.Delete(ctx, getSpecificRouterPath(id))
	return
}

const routersSegment = "/v4/compute/routers"

func getRoutersPath() string {
	return routersSegment
}

func getSpecificRouterPath(routerId int) string {
	return goclient.Join(routersSegment, routerId)
}
