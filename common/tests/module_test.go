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

func TestModuleService(t *testing.T) {
	tests.Handle("/v4/entities/modules", http.MethodGet, tests.StaticResponse(http.StatusOK, `[`+ModuleData+`]`))
	tests.Handle("/v4/entities/modules/{id:\\d+}", http.MethodGet, tests.StaticResponse(http.StatusOK, ModuleData))
	client := tests.Client()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	service := common.NewModuleService(client)

	t.Run("list", func(t *testing.T) {
		modules, err := service.List(ctx, goclient.Cursor{})
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(modules.Items[0], Module) {
			t.Errorf("malformed response: %v", modules.Items[0])
		}
	})

	t.Run("get", func(t *testing.T) {
		module, err := service.Get(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(module, Module) {
			t.Errorf("malformed response: %v", module)
		}
	})
}
