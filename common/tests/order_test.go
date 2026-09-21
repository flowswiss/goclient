package commontests

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/testutil"
)

func TestOrderService(t *testing.T) {
	testutil.Handle("/v4/orders/{id:\\d+}", http.MethodGet, testutil.StaticResponse(http.StatusOK, OrderData))
	client := testutil.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewOrderService(client)

	t.Run("get", func(t *testing.T) {
		order, err := service.Get(ctx, common.OrderGetReq{ID: 1})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(order, Order) {
			t.Errorf("malformed response: %v", order)
		}
	})
}
