package common

import (
	"context"

	"github.com/flowswiss/goclient/v2/core"
)

type Module struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Parent    *Module    `json:"parent"`
	Sorting   int        `json:"sorting"`
	Locations []Location `json:"locations"`
}

type ModuleService struct {
	client *core.Client
}

func NewModuleService(client *core.Client) *ModuleService {
	return &ModuleService{client: client}
}

func (l ModuleService) List(ctx context.Context, cursor core.Cursor) (list List[Module], err error) {
	list.Pagination, err = l.client.List(ctx, getModulesPath(), cursor, &list.Items)
	return
}

type ModuleGetReq struct {
	ID uint `json:"-"`
}

func (l ModuleService) Get(ctx context.Context, req ModuleGetReq) (module Module, err error) {
	err = l.client.Get(ctx, getSpecificModulePath(req.ID), &module)
	return
}

const modulesSegment = "/v4/entities/modules"

func getModulesPath() string {
	return modulesSegment
}

func getSpecificModulePath(moduleID uint) string {
	return core.Join(modulesSegment, moduleID)
}
