package commontests

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
	"github.com/flowswiss/goclient/v2/testutil"
)

func TestProductService(t *testing.T) {
	testutil.Handle("/v4/products", http.MethodGet, testutil.StaticResponse(http.StatusOK, `[`+ProductData+`]`))
	testutil.Handle("/v4/products/{type:[a-z\\-]+}", http.MethodGet, testutil.StaticResponse(http.StatusOK, `[`+ProductData+`]`))
	testutil.Handle("/v4/products/{id:\\d+}", http.MethodGet, testutil.StaticResponse(http.StatusOK, ProductData))
	client := testutil.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewProductService(client)

	t.Run("list", func(t *testing.T) {
		products, err := service.List(ctx, core.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(products.Items[0], Product) {
			t.Errorf("malformed response: %v", products.Items[0])
		}
	})

	t.Run("list-by-type", func(t *testing.T) {
		products, err := service.ListByType(ctx, common.ProductListByTypeReq{
			ProductType: "type",
			Cursor:      core.Cursor{},
		})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(products.Items[0], Product) {
			t.Errorf("malformed response: %v", products.Items[0])
		}
	})

	t.Run("get", func(t *testing.T) {
		product, err := service.Get(ctx, common.ProductGetReq{ID: 1})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(product, Product) {
			t.Errorf("malformed response: %v", product)
		}
	})
}
