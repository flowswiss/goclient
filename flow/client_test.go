package flow

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"testing"
)

func TestNewClientBase(t *testing.T) {
	client := NewClient(nil)
	if client == nil {
		t.Fatal("client should not be null")
	}

	if client.BaseURL.String() != "https://api.flow.swiss/" {
		t.Error("base url is not matching")
	}

	regex, err := regexp.Compile("^flow/\\d+\\.\\d+\\.\\d+$")
	if err != nil {
		t.Fatal(err)
	}

	if !regex.MatchString(client.UserAgent) {
		t.Error("user agent does not match expectation: expected", regex.String(), "got", client.UserAgent)
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient(nil)
	if client == nil {
		t.Fatal("auth should not be null")
	}

	client.BaseURL, _ = url.Parse("https://api.cloudbit.ch/")

	if client.BaseURL.String() != "https://api.cloudbit.ch/" {
		t.Error("base url is not matching")
	}
}

func TestClient_NewRequest(t *testing.T) {
	client := NewClient(nil)

	body := "{\"hello\": \"world\"}"

	buf := &bytes.Buffer{}
	buf.WriteString(body)

	req, err := client.NewRequest(context.Background(), "GET", "/v4/test", buf)
	if err != nil {
		t.Fatal(err)
	}

	if req.Method != "GET" {
		t.Error("expected method GET got", req.Method)
	}

	if req.Host != "api.flow.swiss" {
		t.Error("expected host to be api.flow.swiss got", req.Host)
	}

	if req.URL.String() != "https://api.flow.swiss/v4/test" {
		t.Error("expected url", "https://api.flow.swiss/v4/test", "got", req.URL.String())
	}

	if userAgent := req.Header.Get("User-Agent"); userAgent != client.UserAgent {
		t.Error("expected user agent", client.UserAgent, "got", userAgent)
	}

	if contentType := req.Header.Get("Content-Type"); contentType != "application/json" {
		t.Error("expected content type application/json got", contentType)
	}

	if accept := req.Header.Get("Accept"); accept != "application/json" {
		t.Error("expected accept application/json got", accept)
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(reqBody) != body {
		t.Error("invalid request body")
	}
}

func TestClient_Do(t *testing.T) {
	setupMockServer(t)

	serveMux.HandleFunc("/v4/test", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json")
		res.Header().Add("Link", fmt.Sprintf("%s, %s, %s, %s, %s",
			"<http://localhost/v4/test?page=1&per_page=10>; rel=\"first\"",
			"<http://localhost/v4/test?page=5&per_page=10>; rel=\"last\"",
			"<http://localhost/v4/test?page=2&per_page=10>; rel=\"self\"",
			"<http://localhost/v4/test?page=3&per_page=10>; rel=\"next\"",
			"<http://localhost/v4/test?page=1&per_page=10>; rel=\"prev\"",
		))
		res.Header().Add("X-Pagination-Count", "10")
		res.Header().Add("X-Pagination-Limit", "10")
		res.Header().Add("X-Pagination-Total-Count", "47")
		res.Header().Add("X-Pagination-Current-Page", "2")
		res.Header().Add("X-Pagination-Total-Pages", "5")

		_, err := res.Write([]byte(`{"hello": "world"}`))
		if err != nil {
			t.Fatal(err)
		}
	})

	req, err := client.NewRequest(context.Background(), "GET", "/v4/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	val := &struct {
		Hello string `json:"hello"`
	}{}

	res, err := client.Do(req, val)
	if err != nil {
		t.Fatal(err)
	}

	if val.Hello != "world" {
		t.Error("auth did not parse json response correctly")
	}

	if res.Count != 10 {
		t.Error("auth did not parse pagination correctly:", "expected 10 items got", res.Count)
	}

	if res.Limit != 10 {
		t.Error("auth did not parse pagination correctly:", "expected item limit of 10 got", res.Limit)
	}

	if res.TotalCount != 47 {
		t.Error("auth did not parse pagination correctly:", "expected total item count of 47 got", res.TotalCount)
	}

	if res.CurrentPage != 2 {
		t.Error("auth did not parse pagination correctly:", "expected current page 2 got", res.CurrentPage)
	}

	if res.TotalPages != 5 {
		t.Error("auth did not parse pagination correctly:", "expected total page count of 5 got", res.TotalPages)
	}

	if res.Links.First == "" || res.Links.Last == "" || res.Links.Current == "" || res.Links.Prev == "" || res.Links.Next == "" {
		t.Error("auth did not parse pagination correctly:", "expected all pagination links got", res.Links)
	}
}

func TestClient_DoError(t *testing.T) {
	setupMockServer(t)

	errorMessage := "Oops something went wrong"
	serveMux.HandleFunc("/v4/test", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("X-Request-Id", "0e1a9a390a19ef145717170d381be279bd1afdc83623fd871cb9f020d6a74366")

		body := fmt.Sprintf(`{"error": {"message": {"en": "%s"}}}`, errorMessage)

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500)
		_, err := res.Write([]byte(body))
		if err != nil {
			t.Fatal(err)
		}
	})

	req, err := client.NewRequest(context.Background(), "GET", "/v4/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Do(req, nil)
	if err == nil {
		t.Fatal("expected error but got success")
	}

	errorResponse, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatal("expected error response but got", fmt.Sprintf("%T", err))
	}

	if errorResponse.Message != errorMessage {
		t.Error(fmt.Sprintf("expected error message %q got %q", errorMessage, errorResponse.Message))
	}

	if errorResponse.RequestID == "" {
		t.Error("expected request id in error response")
	}
}

func TestClient_DoContext(t *testing.T) {
	setupMockServer(t)

	ctx, cancel := context.WithCancel(context.Background())

	req, err := client.NewRequest(ctx, "GET", "/v4/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	cancel()

	_, err = client.Do(req, nil)
	if err == nil {
		t.Fatal("expected error but got success")
	}

	urlError, ok := err.(*url.Error)
	if !ok {
		t.Fatal("expected error to be url error but got", fmt.Sprintf("%T", err))
	}

	if urlError.Err != context.Canceled {
		t.Fatal("expected context canceled error but got", urlError.Err)
	}
}

func Test_AddOptions(t *testing.T) {
	base := "/v4/test?q=test"
	options := PaginationOptions{
		Page:    1,
		PerPage: 5,
	}
	expectation := "/v4/test?page=1&per_page=5&q=test"

	res, err := addOptions(base, options)
	if err != nil {
		t.Fatal(err)
	}

	if res != expectation {
		t.Error("expected path", expectation, "got", res)
	}
}
