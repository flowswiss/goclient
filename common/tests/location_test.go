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

func TestLocationService(t *testing.T) {
	tests.Handle("/v4/entities/locations", http.MethodGet, tests.StaticResponse(http.StatusOK, `[`+LocationData+`]`))
	tests.Handle("/v4/entities/locations/{id:\\d+}", http.MethodGet, tests.StaticResponse(http.StatusOK, LocationData))
	client := tests.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewLocationService(client)

	t.Run("list", func(t *testing.T) {
		locations, err := service.List(ctx, goclient.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(locations.Items[0], Location) {
			t.Errorf("malformed response: %v", locations.Items[0])
		}
	})

	t.Run("get", func(t *testing.T) {
		location, err := service.Get(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(location, Location) {
			t.Errorf("malformed response: %v", location)
		}
	})
}
