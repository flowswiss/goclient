package macbaremetal

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

type RouterUpdateReq struct {
	ID uint `json:"-"`

	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (r RouterService) Update(ctx context.Context, req RouterUpdateReq) (router Router, err error) {
	err = r.client.Update(ctx, getSpecificRouterPath(req.ID), req, &router)
	return
}

const routersSegment = "/v4/macbaremetal/routers"

func getRoutersPath() string {
	return routersSegment
}

func getSpecificRouterPath(id uint) string {
	return core.Join(routersSegment, id)
}
