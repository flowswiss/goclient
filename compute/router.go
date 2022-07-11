package compute

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Router struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Location    common.Location `json:"location"`
	Public      bool            `json:"public"`
	SourceNAT   bool            `json:"snat"`
	PublicIP    string          `json:"public_ip"`
}

type RouterList struct {
	Items      []Router
	Pagination goclient.Pagination
}

type RouterCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LocationID  int    `json:"location_id"`
	Public      bool   `json:"public"`
}

type RouterUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Public      bool   `json:"public,omitempty"`
}

type RouterService struct {
	client goclient.Client
}

func NewRouterService(client goclient.Client) RouterService {
	return RouterService{client: client}
}

func (r RouterService) RouterInterfaces(routerID int) RouterInterfaceService {
	return NewRouterInterfaceService(r.client, routerID)
}

func (r RouterService) Routes(routerID int) RouteService {
	return NewRouteService(r.client, routerID)
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

func getSpecificRouterPath(routerID int) string {
	return goclient.Join(routersSegment, routerID)
}
