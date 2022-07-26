package goclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	version   = "2.0.0"
	userAgent = "goclient/" + version
	baseURL   = "https://api.flow.swiss/"
	encoding  = "application/json"
)

var ErrWrongContentType = errors.New("wrong content type")

type APIError struct {
	response  *http.Response
	message   string
	requestID string
}

func (a APIError) Response() *http.Response { return a.response }
func (a APIError) Error() string            { return a.message }
func (a APIError) RequestID() string        { return a.requestID }

type Option func(*Client)

type Client struct {
	httpClient *http.Client
	base       *url.URL
	userAgent  string
}

func NewClient(options ...Option) Client {
	base, _ := url.Parse(baseURL)

	client := Client{
		base:       base,
		userAgent:  userAgent,
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(&client)
	}

	return client
}

func WithBase(base string) Option {
	return func(client *Client) {
		client.base, _ = url.Parse(base)
	}
}

func WithHTTPClientOption(opt func(*http.Client)) Option {
	return func(client *Client) {
		opt(client.httpClient)
	}
}

func WithToken(token string) Option {
	return func(client *Client) {
		client.httpClient.Transport = AuthTransport{
			Token: token,
			Base:  client.httpClient.Transport,
		}
	}
}

func WithUserAgent(userAgent string) Option {
	return func(client *Client) {
		client.userAgent = fmt.Sprintf("%s %s", userAgent, client.userAgent)
	}
}

func (c Client) With(options ...Option) Client {
	for _, option := range options {
		option(&c)
	}
	return c
}

func (c Client) List(ctx context.Context, path string, cursor Cursor, dest interface{}) (Pagination, error) {
	path, err := addOptions(path, cursor)
	if err != nil {
		return Pagination{}, fmt.Errorf("encode query: %w", err)
	}

	req, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return Pagination{}, fmt.Errorf("create list request: %w", err)
	}

	res, err := c.Do(ctx, req, dest)
	if err != nil {
		return Pagination{}, fmt.Errorf("%s %s: %w", req.Method, req.URL.String(), err)
	}

	return ParsePagination(res.Header), nil
}

func (c Client) Get(ctx context.Context, path string, dest interface{}) error {
	req, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return fmt.Errorf("create get request: %w", err)
	}

	_, err = c.Do(ctx, req, dest)
	if err != nil {
		return fmt.Errorf("%s %s: %w", req.Method, req.URL.String(), err)
	}

	return nil
}

func (c Client) Create(ctx context.Context, path string, body interface{}, dest interface{}) error {
	req, err := c.NewRequest(http.MethodPost, path, body)
	if err != nil {
		return fmt.Errorf("create get request: %w", err)
	}

	_, err = c.Do(ctx, req, dest)
	if err != nil {
		return fmt.Errorf("%s %s: %w", req.Method, req.URL.String(), err)
	}

	return nil
}

func (c Client) Update(ctx context.Context, path string, body interface{}, dest interface{}) error {
	req, err := c.NewRequest(http.MethodPatch, path, body)
	if err != nil {
		return fmt.Errorf("create get request: %w", err)
	}

	_, err = c.Do(ctx, req, dest)
	if err != nil {
		return fmt.Errorf("%s %s: %w", req.Method, req.URL.String(), err)
	}

	return nil
}

func (c Client) Set(ctx context.Context, path string, body interface{}, dest interface{}) error {
	req, err := c.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return fmt.Errorf("create get request: %w", err)
	}

	_, err = c.Do(ctx, req, dest)
	if err != nil {
		return fmt.Errorf("%s %s: %w", req.Method, req.URL.String(), err)
	}

	return nil
}

func (c Client) Delete(ctx context.Context, path string) error {
	req, err := c.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return fmt.Errorf("create delete request: %w", err)
	}

	_, err = c.Do(ctx, req, nil)
	if err != nil {
		return fmt.Errorf("%s %s: %w", req.Method, req.URL.String(), err)
	}

	return nil
}

func (c Client) NewRequest(method string, path string, body interface{}) (*http.Request, error) {
	u, err := c.base.Parse(path)
	if err != nil {
		return nil, err
	}

	reader, ok := body.(io.Reader)
	if !ok {
		buf := &bytes.Buffer{}
		if body != nil {
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}
		}
		reader = buf
	}

	req, err := http.NewRequest(method, u.String(), reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", encoding)
	req.Header.Add("Accept", encoding)
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, val interface{}) (*http.Response, error) {
	res, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent {
		return res, nil
	}

	if res.StatusCode >= 400 {
		apiError := &struct {
			Error struct {
				Message struct {
					En string `json:"en"`
				} `json:"message"`
			} `json:"error"`
		}{}

		err := json.NewDecoder(res.Body).Decode(apiError)
		if err != nil {
			return nil, fmt.Errorf("parse response body: %w", err)
		}

		return nil, APIError{
			response:  res,
			message:   apiError.Error.Message.En,
			requestID: res.Header.Get("X-Request-Id"),
		}
	}

	if writer, ok := val.(io.Writer); ok {
		_, err = io.Copy(writer, res.Body)
		if err != nil {
			return nil, fmt.Errorf("read response body: %w", err)
		}
	} else if val != nil {
		err = json.NewDecoder(res.Body).Decode(val)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return res, nil
			}

			return nil, fmt.Errorf("parse response body: %w", err)
		}
	}

	return res, nil
}

func addOptions(path string, options interface{}) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	newQuery, err := query.Values(options)
	if err != nil {
		return "", err
	}

	prevQuery := u.Query()
	for key, arr := range newQuery {
		for _, val := range arr {
			prevQuery.Add(key, val)
		}
	}
	u.RawQuery = prevQuery.Encode()

	return u.String(), nil
}
