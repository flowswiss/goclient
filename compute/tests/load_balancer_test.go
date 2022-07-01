package computetests

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/flowswiss/goclient"
	commontests "github.com/flowswiss/goclient/common/tests"
	"github.com/flowswiss/goclient/compute"
	"github.com/flowswiss/goclient/internal/tests"
)

func TestLoadBalancerService(t *testing.T) {
	tests.Handle("/v4/compute/load-balancers", http.MethodGet, tests.StaticResponse(http.StatusOK, `[`+LoadBalancerData+`]`))
	tests.Handle("/v4/compute/load-balancers", http.MethodPost, tests.StaticResponse(http.StatusCreated, commontests.OrderingData))
	tests.Handle("/v4/compute/load-balancers/{id:\\d+}", http.MethodGet, tests.StaticResponse(http.StatusOK, LoadBalancerData))
	tests.Handle("/v4/compute/load-balancers/{id:\\d+}", http.MethodPatch, tests.StaticResponse(http.StatusOK, LoadBalancerData))
	tests.Handle("/v4/compute/load-balancers/{id:\\d+}", http.MethodDelete, tests.StaticResponse(http.StatusNoContent, ``))
	client := tests.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := compute.NewLoadBalancerService(client)

	t.Run("list", func(t *testing.T) {
		loadBalancers, err := service.List(ctx, goclient.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(loadBalancers.Items[0], LoadBalancer) {
			t.Errorf("malformed response: %v", loadBalancers.Items[0])
		}
	})

	t.Run("create", func(t *testing.T) {
		ordering, err := service.Create(ctx, compute.LoadBalancerCreate{
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
		loadBalancer, err := service.Get(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(loadBalancer, LoadBalancer) {
			t.Errorf("malformed response: %v", loadBalancer)
		}
	})

	t.Run("update", func(t *testing.T) {
		loadBalancer, err := service.Update(ctx, 1, compute.LoadBalancerUpdate{
			Name: "lb-test",
		})

		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(loadBalancer, LoadBalancer) {
			t.Errorf("malformed response: %v", loadBalancer)
		}
	})

	t.Run("delete", func(t *testing.T) {
		err := service.Delete(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}
	})
}
