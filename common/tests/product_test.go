package commontests

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
	"github.com/flowswiss/goclient/internal/tests"
)

func TestProductService(t *testing.T) {
	tests.Handle("/v4/products", http.MethodGet, tests.StaticResponse(http.StatusOK, `[`+ProductData+`]`))
	tests.Handle("/v4/products/{type:[a-z\\-]+}", http.MethodGet, tests.StaticResponse(http.StatusOK, `[`+ProductData+`]`))
	tests.Handle("/v4/products/{id:\\d+}", http.MethodGet, tests.StaticResponse(http.StatusOK, ProductData))
	client := tests.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewProductService(client)

	t.Run("list", func(t *testing.T) {
		products, err := service.List(ctx, goclient.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(products.Items[0], Product) {
			t.Errorf("malformed response: %v", products.Items[0])
		}
	})

	t.Run("list-by-type", func(t *testing.T) {
		products, err := service.ListByType(ctx, "type", goclient.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(products.Items[0], Product) {
			t.Errorf("malformed response: %v", products.Items[0])
		}
	})

	t.Run("get", func(t *testing.T) {
		product, err := service.Get(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(product, Product) {
			t.Errorf("malformed response: %v", product)
		}
	})
}
