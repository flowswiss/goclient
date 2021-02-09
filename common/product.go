package common

import (
	"context"

	"github.com/flowswiss/goclient"
)

type ProductType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ProductUsageCycle struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

type ProductItem struct {
	Id          int    `json:"id"`
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
	Id             int                   `json:"id"`
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

type ProductList struct {
	Items      []Product
	Pagination goclient.Pagination
}

type ProductService struct {
	client goclient.Client
}

func NewProductService(client goclient.Client) ProductService {
	return ProductService{client: client}
}

func (p ProductService) List(ctx context.Context, cursor goclient.Cursor) (list ProductList, err error) {
	list.Pagination, err = p.client.List(ctx, getProductsPath(), cursor, &list.Items)
	return
}

func (p ProductService) ListByType(ctx context.Context, productType string, cursor goclient.Cursor) (list ProductList, err error) {
	list.Pagination, err = p.client.List(ctx, getProductsByTypePath(productType), cursor, &list.Items)
	return
}

func (p ProductService) Get(ctx context.Context, id int) (product Product, err error) {
	err = p.client.Get(ctx, getSpecificProductPath(id), &product)
	return
}

const productsSegment = "/v4/products"

func getProductsPath() string {
	return productsSegment
}

func getProductsByTypePath(productType string) string {
	return goclient.Join(productsSegment, productType)
}

func getSpecificProductPath(id int) string {
	return goclient.Join(productsSegment, id)
}
