package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type RouterService struct {
	client *core.Client
}

func NewRouterService(client *core.Client) *RouterService {
	return &RouterService{client: client}
}

func (r RouterService) List(ctx context.Context, cursor core.Cursor) (list common.List[Router], err error) {
	list.Pagination, err = r.client.List(ctx, getRoutersPath(), cursor, &list.Items)
	return
}

type RouterGetReq struct {
	ID uint `json:"-"`
}

func (r RouterService) Get(ctx context.Context, req RouterGetReq) (router Router, err error) {
	err = r.client.Get(ctx, getSpecificRouterPath(req.ID), &router)
	return
}

type RouterCreateReq struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	LocationID  int     `json:"location_id"`
	Public      bool    `json:"public"`
}

func (r RouterService) Create(ctx context.Context, req RouterCreateReq) (router Router, err error) {
	err = r.client.Create(ctx, getRoutersPath(), req, &router)
	return
}

type RouterUpdateReq struct {
	ID uint `json:"-"`

	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Public      *bool   `json:"public,omitempty"`
}

func (r RouterService) Update(ctx context.Context, req RouterUpdateReq) (router Router, err error) {
	err = r.client.Update(ctx, getSpecificRouterPath(req.ID), req, &router)
	return
}

type RouterDeleteReq struct {
	ID uint `json:"-"`
}

func (r RouterService) Delete(ctx context.Context, req RouterDeleteReq) (err error) {
	err = r.client.Delete(ctx, getSpecificRouterPath(req.ID))
	return
}

const routersSegment = "/v4/compute/routers"

func getRoutersPath() string {
	return routersSegment
}

func getSpecificRouterPath(routerID uint) string {
	return core.Join(routersSegment, routerID)
}
