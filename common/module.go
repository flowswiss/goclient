package common

import (
	"context"

	"github.com/flowswiss/goclient"
)

type Module struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Parent    *Module    `json:"parent"`
	Sorting   int        `json:"sorting"`
	Locations []Location `json:"locations"`
}

type ModuleList struct {
	goclient.Pagination
	Items []Module
}

type ModuleService struct {
	client goclient.Client
}

func NewModuleService(client goclient.Client) ModuleService {
	return ModuleService{client: client}
}

func (l ModuleService) List(ctx context.Context, cursor goclient.Cursor) (list ModuleList, err error) {
	list.Pagination, err = l.client.List(ctx, getModulesPath(), cursor, &list.Items)
	return
}

func (l ModuleService) Get(ctx context.Context, id int) (module Module, err error) {
	err = l.client.Get(ctx, getSpecificModulePath(id), &module)
	return
}

const modulesSegment = "/v4/entities/modules"

func getModulesPath() string {
	return modulesSegment
}

func getSpecificModulePath(moduleID int) string {
	return goclient.Join(modulesSegment, moduleID)
}
