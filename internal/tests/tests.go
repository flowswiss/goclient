package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"

	"github.com/flowswiss/goclient"
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

func Client() goclient.Client {
	server := httptest.NewServer(router)

	return goclient.NewClient(
		goclient.WithBase(server.URL),
	)
}
