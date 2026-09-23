package computetests

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	commontests "github.com/flowswiss/goclient/v2/common/tests"
	"github.com/flowswiss/goclient/v2/compute"
	"github.com/flowswiss/goclient/v2/core"
	"github.com/flowswiss/goclient/v2/testutil"
)

func TestLoadBalancerService(t *testing.T) {
	testutil.Handle("/v4/compute/load-balancers", http.MethodGet, testutil.StaticResponse(http.StatusOK, `[`+LoadBalancerData+`]`))
	testutil.Handle("/v4/compute/load-balancers", http.MethodPost, testutil.StaticResponse(http.StatusCreated, commontests.OrderingData))
	testutil.Handle("/v4/compute/load-balancers/{id:\\d+}", http.MethodGet, testutil.StaticResponse(http.StatusOK, LoadBalancerData))
	testutil.Handle("/v4/compute/load-balancers/{id:\\d+}", http.MethodPatch, testutil.StaticResponse(http.StatusOK, LoadBalancerData))
	testutil.Handle("/v4/compute/load-balancers/{id:\\d+}", http.MethodDelete, testutil.StaticResponse(http.StatusNoContent, ``))
	client := testutil.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := compute.NewLoadBalancerService(client)

	t.Run("list", func(t *testing.T) {
		loadBalancers, err := service.List(ctx, core.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(loadBalancers.Items[0], LoadBalancer) {
			t.Errorf("malformed response: %v", loadBalancers.Items[0])
		}
	})

	t.Run("create", func(t *testing.T) {
		ordering, err := service.Create(ctx, compute.LoadBalancerCreateReq{
			Name:             "lb-test",
			LocationID:       1,
			AttachExternalIP: true,
			NetworkID:        0,
			PrivateIP:        "172.0.0.1",
		})

		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(ordering, commontests.Ordering) {
			t.Errorf("malformed response: %v", ordering)
		}
	})

	t.Run("get", func(t *testing.T) {
		loadBalancer, err := service.Get(ctx, compute.LoadBalancerGetReq{ID: 1})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(loadBalancer, LoadBalancer) {
			t.Errorf("malformed response: %v", loadBalancer)
		}
	})

	t.Run("update", func(t *testing.T) {
		loadBalancer, err := service.Update(ctx, compute.LoadBalancerUpdateReq{
			Name: new("lb-test"),
		})

		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(loadBalancer, LoadBalancer) {
			t.Errorf("malformed response: %v", loadBalancer)
		}
	})

	t.Run("delete", func(t *testing.T) {
		err := service.Delete(ctx, compute.LoadBalancerDeleteReq{ID: 1})
		if err != nil {
			t.Fatal(err)
		}
	})
}
