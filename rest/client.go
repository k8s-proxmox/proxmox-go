package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"golang.org/x/time/rate"
)

const (
	defaultUserAgent = "sp-yduck/proxmox-go"
	defaultQPS       = 20
)

type RESTClient struct {
	// proxmox rest api endpoint
	endpoint string

	httpClient *http.Client

	transport Transport

	rateLimiter *rate.Limiter

	logger logr.Logger
}

type TicketRequest struct {
	// required
	Username string `json:"username"`
	Password string `json:"password"`
	// optional
	Otp   string `json:"otp,omitempty"`
	Path  string `json:"path,omitempty"`
	Privs string `json:"privs,omitempty"`
	Realm string `json:"realm,omitempty"`
}

type ClientOption func(*RESTClient) error

func NewRESTClient(baseUrl string, opts ...ClientOption) (*RESTClient, error) {
	client := &RESTClient{
		endpoint:    complementURL(baseUrl),
		httpClient:  &http.Client{},
		rateLimiter: rate.NewLimiter(rate.Every(1*time.Second), defaultQPS),
	}

	for _, option := range opts {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func complementURL(url string) string {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	url, _ = strings.CutSuffix(url, "/")
	return url
}

func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c *RESTClient) error {
		c.transport.SetBase(transport)
		return nil
	}
}

func WithUserPassword(username, password string) ClientOption {
	return func(c *RESTClient) error {
		var err error
		c.transport.AuthProvider, err = NewTicketProvider(c.endpoint, username, password)
		if err != nil {
			return err
		}
		return nil
	}
}

func WithAPIToken(tokenid, secret string) ClientOption {
	return func(c *RESTClient) error {
		c.transport.AuthProvider = NewTokenProvider(tokenid, secret)
		return nil
	}
}

func WithQPS(qps int) ClientOption {
	return func(c *RESTClient) error {
		c.rateLimiter = rate.NewLimiter(rate.Every(1*time.Second), qps)
		return nil
	}
}

func WithLogger(logger logr.Logger) ClientOption {
	return func(c *RESTClient) error {
		c.logger = logger
		return nil
	}
}

func (c *RESTClient) SetMaxQPS(qps int) {
	c.rateLimiter = rate.NewLimiter(rate.Every(1*time.Second), qps)
}

func (c *RESTClient) Do(ctx context.Context, httpMethod, urlPath string, req, v interface{}) error {
	endpoint := c.endpoint + urlPath
	c.logger.V(1).Info(fmt.Sprintf("making %s request for %s", httpMethod, endpoint))

	var body io.Reader
	if req != nil {
		jsonReq, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = bytes.NewReader(jsonReq)
		c.logger.WithValues("endpoint", endpoint).V(1).Info(fmt.Sprintf("request body: %s", string(jsonReq)))
	}

	httpReq, err := http.NewRequestWithContext(ctx, httpMethod, endpoint, body)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Accept", "application/json")

	if err := c.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpRsp.Body.Close()

	buf, err := checkResponse(httpRsp)
	if err != nil {
		c.logger.V(0).Error(err, fmt.Sprintf("responce of %s:(%s): %s", endpoint, httpMethod, string(buf)))
		return err
	}
	c.logger.V(1).Info(fmt.Sprintf("responce of %s:(%s): %s", endpoint, httpMethod, string(buf)))

	// try unmarshal on {"data": any} firstly
	var datakey map[string]json.RawMessage
	if err := json.Unmarshal(buf, &datakey); err != nil {
		return err
	}
	if body, ok := datakey["data"]; ok {
		return json.Unmarshal(body, &v)
	}

	return json.Unmarshal(buf, &v)
}

func (c *RESTClient) Get(ctx context.Context, path string, res interface{}) error {
	return c.Do(ctx, http.MethodGet, path, nil, res)
}

func (c *RESTClient) Post(ctx context.Context, path string, req, res interface{}) error {
	return c.Do(ctx, http.MethodPost, path, req, res)
}

func (c *RESTClient) Put(ctx context.Context, path string, req, res interface{}) error {
	return c.Do(ctx, http.MethodPut, path, req, res)
}

func (c *RESTClient) Delete(ctx context.Context, path string, req, res interface{}) error {
	return c.Do(ctx, http.MethodDelete, path, req, res)
}

func checkResponse(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body while handling http response of status %d : %v", res.StatusCode, err)
	}

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return body, nil
	}

	return nil, NewError(res.StatusCode, res.Status, body)
}
