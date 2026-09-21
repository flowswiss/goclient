package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type RouterInterfaceService struct {
	client *core.Client
}

func NewRouterInterfaceService(client *core.Client) *RouterInterfaceService {
	return &RouterInterfaceService{client: client}
}

type RouterInterfaceListReq struct {
	RouterID uint `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (r RouterInterfaceService) List(ctx context.Context, req RouterInterfaceListReq) (
	list common.List[RouterInterface],
	err error,
) {
	list.Pagination, err = r.client.List(ctx, getRouterInterfacesPath(req.RouterID), req.Cursor, &list.Items)
	return
}

const routerInterfacesSegment = "router-interfaces"

func getRouterInterfacesPath(id uint) string {
	return core.Join(routersSegment, id, routerInterfacesSegment)
}
