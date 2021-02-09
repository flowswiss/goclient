package compute

import (
	"context"

	"github.com/flowswiss/goclient"
)

type RouterInterface struct {
	Id        int     `json:"id"`
	PrivateIp string  `json:"private_ip"`
	Network   Network `json:"network"`
}

type RouterInterfaceList struct {
	Items      []RouterInterface
	Pagination goclient.Pagination
}

type RouterInterfaceCreate struct {
	NetworkId int    `json:"network_id"`
	PrivateIp string `json:"private_ip,omitempty"`
}

type RouterInterfaceService struct {
	client   goclient.Client
	routerId int
}

func NewRouterInterfaceService(client goclient.Client, routerId int) RouterInterfaceService {
	return RouterInterfaceService{client: client, routerId: routerId}
}

func (r RouterInterfaceService) List(ctx context.Context, cursor goclient.Cursor) (list RouterInterfaceList, err error) {
	list.Pagination, err = r.client.List(ctx, getRouterInterfacesPath(r.routerId), cursor, &list.Items)
	return
}

func (r RouterInterfaceService) Create(ctx context.Context, body RouterInterfaceCreate) (routerInterface RouterInterface, err error) {
	err = r.client.Create(ctx, getRouterInterfacesPath(r.routerId), body, &routerInterface)
	return
}

func (r RouterInterfaceService) Delete(ctx context.Context, id int) (err error) {
	err = r.client.Delete(ctx, getSpecificRouterInterfacePath(r.routerId, id))
	return
}

const routerInterfacesSegment = "router-interfaces"

func getRouterInterfacesPath(routerId int) string {
	return goclient.Join(routersSegment, routerId, routerInterfacesSegment)
}

func getSpecificRouterInterfacePath(routerId, routerInterfaceId int) string {
	return goclient.Join(routersSegment, routerId, routerInterfacesSegment, routerInterfaceId)
}
