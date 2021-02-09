package compute

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type ElasticIpProduct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ElasticIpAttachment struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ElasticIp struct {
	Id         int                 `json:"id"`
	Product    ElasticIpProduct    `json:"product"`
	Location   common.Location     `json:"location"`
	Price      float64             `json:"price"`
	PublicIp   string              `json:"public_ip"`
	PrivateIp  string              `json:"private_ip"`
	Attachment ElasticIpAttachment `json:"attached_instance"`
}

type ElasticIpList struct {
	Items      []ElasticIp
	Pagination goclient.Pagination
}

type ElasticIpCreate struct {
	LocationId int `json:"location_id"`
}

type ElasticIpService struct {
	client goclient.Client
}

func NewElasticIpService(client goclient.Client) ElasticIpService {
	return ElasticIpService{client: client}
}

func (e ElasticIpService) List(ctx context.Context, cursor goclient.Cursor) (list ElasticIpList, err error) {
	list.Pagination, err = e.client.List(ctx, getElasticIpsPath(), cursor, &list.Items)
	return
}

func (e ElasticIpService) Create(ctx context.Context, body ElasticIpCreate) (elasticIp ElasticIp, err error) {
	err = e.client.Create(ctx, getElasticIpsPath(), body, &elasticIp)
	return
}

func (e ElasticIpService) Delete(ctx context.Context, id int) (err error) {
	err = e.client.Delete(ctx, getSpecificElasticIpPath(id))
	return
}

const elasticIpsSegment = "/v4/computes/elastic-ips"

func getElasticIpsPath() string {
	return elasticIpsSegment
}

func getSpecificElasticIpPath(elasticIpId int) string {
	return goclient.Join(elasticIpsSegment, elasticIpId)
}
