package common

import (
	"context"

	"github.com/flowswiss/goclient/v2/core"
)

type ProductType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ProductUsageCycle struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

type ProductItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type ProductAvailability struct {
	Location  Location `json:"location"`
	Available int      `json:"available"`
}

type DeploymentFee struct {
	Location        Location `json:"location"`
	Price           float64  `json:"price"`
	FreeDeployments int      `json:"free_deployments"`
}

type Product struct {
	ID             int                   `json:"id"`
	Name           string                `json:"product_name"`
	Type           ProductType           `json:"type"`
	Visibility     string                `json:"visibility"`
	UsageCycle     ProductUsageCycle     `json:"usage_cycle"`
	Items          []ProductItem         `json:"items"`
	Price          float64               `json:"price"`
	Availability   []ProductAvailability `json:"availability"`
	Category       string                `json:"category"`
	DeploymentFees []DeploymentFee       `json:"deployment_fees"`
}

type BriefProduct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ProductService struct {
	client *core.Client
}

func NewProductService(client *core.Client) *ProductService {
	return &ProductService{client: client}
}

func (p ProductService) List(ctx context.Context, cursor core.Cursor) (list List[Product], err error) {
	list.Pagination, err = p.client.List(ctx, getProductsPath(), cursor, &list.Items)
	return
}

type ProductListByTypeReq struct {
	ProductType string `json:"-"`

	Cursor core.Cursor `json:"-"`
}

func (p ProductService) ListByType(ctx context.Context, req ProductListByTypeReq) (list List[Product], err error) {
	list.Pagination, err = p.client.List(ctx, getProductsByTypePath(req.ProductType), req.Cursor, &list.Items)
	return
}

type ProductGetReq struct {
	ID uint `json:"-"`
}

func (p ProductService) Get(ctx context.Context, req ProductGetReq) (product Product, err error) {
	err = p.client.Get(ctx, getSpecificProductPath(req.ID), &product)
	return
}

func (p ProductService) ListTypes(ctx context.Context, cursor core.Cursor) (list List[ProductType], err error) {
	list.Pagination, err = p.client.List(ctx, getProductTypesPath(), cursor, &list.Items)
	return
}

const (
	productsSegment     = "/v4/products"
	productTypesSegment = "/v4/entities/product-types"
)

func getProductsPath() string {
	return productsSegment
}

func getProductsByTypePath(productType string) string {
	return core.Join(productsSegment, productType)
}

func getSpecificProductPath(id uint) string {
	return core.Join(productsSegment, id)
}

func getProductTypesPath() string {
	return productTypesSegment
}
