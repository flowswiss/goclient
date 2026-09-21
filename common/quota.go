package common

import (
	"context"
	"fmt"
	"net/url"

	"github.com/flowswiss/goclient/v2/core"
)

type Quota struct {
	Scope  QuotaScope   `json:"scope"`
	Usages []QuotaUsage `json:"usages"`
}

type QuotaScope struct {
	Current  string   `json:"current"`
	Parent   string   `json:"parent"`
	Children []string `json:"children"`
}

type QuotaUsage struct {
	Entity     QuotaEntity `json:"entity"`
	EntityType QuotaEntity `json:"entity_type"`
	Unit       string      `json:"unit"`
	Quota      int         `json:"quota"`
	Current    int         `json:"current"`
}

type QuotaEntity struct {
	ID   int    `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

type QuotaService struct {
	client *core.Client
}

func NewQuotaService(client *core.Client) *QuotaService {
	return &QuotaService{client: client}
}

type QuotaGetReq struct {
	Scope string `json:"-"`
}

func (q QuotaService) Get(ctx context.Context, req QuotaGetReq) (quotas []Quota, err error) {
	query := url.Values{
		"scope": []string{req.Scope},
	}

	path := fmt.Sprint(getQuotasPath(), "?", query.Encode())
	err = q.client.Get(ctx, path, &quotas)
	return
}

const quotasSegment = "/v4/quotas"

func getQuotasPath() string {
	return quotasSegment
}
