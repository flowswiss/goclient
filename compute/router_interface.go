package compute

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

type RouterInterfaceCreateReq struct {
	RouterID uint `json:"-"`

	NetworkID int    `json:"network_id"`
	PrivateIP string `json:"private_ip,omitempty"`
}

func (r RouterInterfaceService) Create(
	ctx context.Context,
	req RouterInterfaceCreateReq,
) (routerInterface RouterInterface, err error) {
	err = r.client.Create(ctx, getRouterInterfacesPath(req.RouterID), req, &routerInterface)
	return
}

type RouterInterfaceDeleteReq struct {
	RouterID          uint `json:"-"`
	RouterInterfaceID uint `json:"-"`
}

func (r RouterInterfaceService) Delete(ctx context.Context, req RouterInterfaceDeleteReq) (err error) {
	err = r.client.Delete(ctx, getSpecificRouterInterfacePath(req.RouterID, req.RouterInterfaceID))
	return
}

const routerInterfacesSegment = "interfaces"

func getRouterInterfacesPath(routerID uint) string {
	return core.Join(routersSegment, routerID, routerInterfacesSegment)
}

func getSpecificRouterInterfacePath(routerID, routerInterfaceID uint) string {
	return core.Join(routersSegment, routerID, routerInterfacesSegment, routerInterfaceID)
}
