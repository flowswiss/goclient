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

func TestModuleService(t *testing.T) {
	testutil.Handle("/v4/entities/modules", http.MethodGet, testutil.StaticResponse(http.StatusOK, `[`+ModuleData+`]`))
	testutil.Handle("/v4/entities/modules/{id:\\d+}", http.MethodGet, testutil.StaticResponse(http.StatusOK, ModuleData))
	client := testutil.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewModuleService(client)

	t.Run("list", func(t *testing.T) {
		modules, err := service.List(ctx, core.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(modules.Items[0], Module) {
			t.Errorf("malformed response: %v", modules.Items[0])
		}
	})

	t.Run("get", func(t *testing.T) {
		module, err := service.Get(ctx, common.ModuleGetReq{ID: 1})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(module, Module) {
			t.Errorf("malformed response: %v", module)
		}
	})
}
