package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/flowswiss/goclient/v2/core"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func StaticResponse(status int, data string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(status)
		_, _ = fmt.Fprint(res, data)
	}
}

func Handle(pattern string, method string, handler http.Handler) {
	router.Path(pattern).Methods(method).Handler(handler)
}

func Client() *core.Client {
	server := httptest.NewServer(router)

	baseURL, _ := url.Parse(server.URL)

	return core.NewClient(core.ClientOpts{
		BaseURL: baseURL,
	})
}
