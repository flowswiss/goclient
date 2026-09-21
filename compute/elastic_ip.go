package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type ElasticIPService struct {
	client *core.Client
}

func NewElasticIPService(client *core.Client) *ElasticIPService {
	return &ElasticIPService{client: client}
}

func (e ElasticIPService) List(ctx context.Context, cursor core.Cursor) (list common.List[ElasticIP], err error) {
	list.Pagination, err = e.client.List(ctx, getElasticIPsPath(), cursor, &list.Items)
	return
}

type ElasticIPCreateReq struct {
	LocationID int `json:"location_id"`
}

func (e ElasticIPService) Create(ctx context.Context, req ElasticIPCreateReq) (elasticIP ElasticIP, err error) {
	err = e.client.Create(ctx, getElasticIPsPath(), req, &elasticIP)
	return
}

type ElasticIPDeleteReq struct {
	ID int `json:"-"`
}

func (e ElasticIPService) Delete(ctx context.Context, req ElasticIPDeleteReq) (err error) {
	err = e.client.Delete(ctx, getSpecificElasticIPPath(req.ID))
	return
}

const elasticIPsSegment = "/v4/compute/elastic-ips"

func getElasticIPsPath() string {
	return elasticIPsSegment
}

func getSpecificElasticIPPath(elasticIPID int) string {
	return core.Join(elasticIPsSegment, elasticIPID)
}
