package macbaremetal

import (
	"context"

	"github.com/flowswiss/goclient"
)

type RouterInterface struct {
	ID        int     `json:"id"`
	PrivateIP string  `json:"private_ip"`
	Network   Network `json:"network"`
}

type RouterInterfaceList struct {
	Items      []RouterInterface
	Pagination goclient.Pagination
}

type RouterInterfaceService struct {
	client   goclient.Client
	routerID int
}

func NewRouterInterfaceService(client goclient.Client, routerID int) RouterInterfaceService {
	return RouterInterfaceService{client: client, routerID: routerID}
}

func (r RouterInterfaceService) List(ctx context.Context, cursor goclient.Cursor) (list RouterInterfaceList, err error) {
	list.Pagination, err = r.client.List(ctx, getRouterInterfacesPath(r.routerID), cursor, &list.Items)
	return
}

const routerInterfacesSegment = "router-interfaces"

func getRouterInterfacesPath(id int) string {
	return goclient.Join(routersSegment, id, routerInterfacesSegment)
}
