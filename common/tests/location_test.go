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

func TestLocationService(t *testing.T) {
	testutil.Handle("/v4/entities/locations", http.MethodGet, testutil.StaticResponse(http.StatusOK, `[`+LocationData+`]`))
	testutil.Handle("/v4/entities/locations/{id:\\d+}", http.MethodGet, testutil.StaticResponse(http.StatusOK, LocationData))
	client := testutil.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewLocationService(client)

	t.Run("list", func(t *testing.T) {
		locations, err := service.List(ctx, core.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(locations.Items[0], Location) {
			t.Errorf("malformed response: %v", locations.Items[0])
		}
	})

	t.Run("get", func(t *testing.T) {
		location, err := service.Get(ctx, common.LocationGetReq{ID: 1})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(location, Location) {
			t.Errorf("malformed response: %v", location)
		}
	})
}
