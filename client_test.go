package goclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func errorMessage(message string) string {
	return `{"error": {"message": {"en": "` + message + `"}}}`
}

func TestClient_Do(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	type data struct {
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	}

	item := data{
		Name:  "Item",
		Value: 12,
	}

	router := mux.NewRouter()
	router.Path("/list").Methods(http.MethodGet).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		res.Header().Set("X-Pagination-Current-Page", query.Get("page"))
		res.Header().Set("X-Pagination-Limit", query.Get("per_page"))
		res.WriteHeader(http.StatusOK)
	})
	router.Path("/get").Methods(http.MethodGet).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(res).Encode(item)
	})
	router.Path("/create").Methods(http.MethodPost).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var val data

		err := json.NewDecoder(req.Body).Decode(&val)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(res, errorMessage(err.Error()))
			return
		}

		res.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(res).Encode(val)
	})
	router.Path("/update").Methods(http.MethodPatch).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var val data

		err := json.NewDecoder(req.Body).Decode(&val)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(res, errorMessage(err.Error()))
			return
		}

		res.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(res).Encode(val)
	})
	router.Path("/delete").Methods(http.MethodDelete).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNoContent)
	})

	server := httptest.NewServer(router)

	client := NewClient(WithBase(server.URL))

	t.Run("list", func(t *testing.T) {
		cursor := Cursor{PerPage: 5, Page: 2}

		pagination, err := client.List(ctx, "/list", cursor, nil)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(pagination.Cursor, cursor) {
			t.Errorf("expected cursor to equal %v, got %v", cursor, pagination.Cursor)
		}
	})

	t.Run("get", func(t *testing.T) {
		var res data
		err := client.Get(ctx, "/get", &res)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(item, res) {
			t.Errorf("expected response to equal %v, got %v", item, res)
		}
	})

	t.Run("create", func(t *testing.T) {
		val := data{
			Name:  "Test create",
			Value: 0.5867,
		}

		var res data
		err := client.Create(ctx, "/create", val, &res)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(val, res) {
			t.Errorf("expected response to equal %v, got %v", val, res)
		}
	})

	t.Run("update", func(t *testing.T) {
		val := data{
			Name:  "Test update",
			Value: 3.1415926,
		}

		var res data
		err := client.Update(ctx, "/update", val, &res)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(val, res) {
			t.Errorf("expected response to equal %v, got %v", val, res)
		}
	})

	t.Run("delete", func(t *testing.T) {
		err := client.Delete(ctx, "/delete")
		if err != nil {
			t.Fatal(err)
		}
	})
}
