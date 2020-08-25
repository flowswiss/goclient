package flow

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/flowswiss/goclient/flow/auth"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
)

const (
	version   = "1.0.0"
	userAgent = "flow/" + version
	baseUrl   = "https://api.flow.swiss/"
	encoding  = "application/json"
)

type Id uint

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	client *http.Client

	// General entities
	Product  ProductService
	Location LocationService
	Module   ModuleService
	Image    ImageService

	// Compute
	Server           ServerService
	ServerAttachment ServerAttachmentService
	KeyPair          KeyPairService
	Network          NetworkService
	ElasticIp        ElasticIpService

	// Other
	Order OrderService
}

type Response struct {
	*http.Response
	Pagination
}

func NewClientWithToken(token string) *Client {
	return NewClient(auth.NewClientWithToken(token))
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(baseUrl)

	client := &Client{
		BaseURL:   baseUrl,
		UserAgent: userAgent,
		client:    httpClient,
	}

	client.Product = &productService{client}
	client.Location = &locationService{client}
	client.Module = &moduleService{client}
	client.Image = &imageService{client}
	client.Server = &serverService{client}
	client.ServerAttachment = &serverAttachmentService{client}
	client.KeyPair = &keyPairService{client}
	client.Network = &networkService{client}
	client.ElasticIp = &elasticIpService{client}
	client.Order = &orderService{client}

	return client
}

func (c *Client) NewRequest(ctx context.Context, method string, path string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(path)
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

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", encoding)
	req.Header.Add("Accept", encoding)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request, val interface{}) (*Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &Response{
		Response:   res,
		Pagination: parsePagination(res),
	}

	if res.StatusCode == http.StatusNoContent {
		return response, nil
	}

	if res.Header.Get("Content-Type") != encoding {
		return nil, ErrorUnsupportedContentType
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
			return nil, err
		}

		return nil, &ErrorResponse{
			Response:  res,
			Message:   apiError.Error.Message.En,
			RequestID: res.Header.Get("X-Request-Id"),
		}
	}

	if writer, ok := val.(io.Writer); ok {
		_, err := io.Copy(writer, res.Body)
		if err != nil {
			return nil, err
		}
	} else if val != nil {
		err := json.NewDecoder(res.Body).Decode(val)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
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
