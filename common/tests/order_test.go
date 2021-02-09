package commontests

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/flowswiss/goclient/common"
	"github.com/flowswiss/goclient/internal/tests"
)

func TestOrderService(t *testing.T) {
	tests.Handle("/v4/orders/{id:\\d+}", http.MethodGet, tests.StaticResponse(http.StatusOK, OrderData))
	client := tests.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewOrderService(client)

	t.Run("get", func(t *testing.T) {
		order, err := service.Get(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(order, Order) {
			t.Errorf("malformed response: %v", order)
		}
	})
}
