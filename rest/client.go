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

	"github.com/sp-yduck/proxmox-go/api"
)

const (
	defaultUserAgent = "sp-yduck/proxmox-go"
)

type RESTClient struct {
	endpoint    string
	httpClient  *http.Client
	tokenid     string
	token       string
	session     *api.Session
	credentials *TicketRequest
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

type ClientOption func(*RESTClient)

func NewRESTClient(baseUrl string, opts ...ClientOption) (*RESTClient, error) {
	client := &RESTClient{
		endpoint:   complementURL(baseUrl),
		httpClient: &http.Client{},
	}

	for _, option := range opts {
		option(client)
	}

	if client.token == "" && client.session == nil && client.credentials != nil {
		ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Minute)
		defer cancel()
		if err := client.makeNewSession(ctx); err != nil {
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

func WithClient(client *http.Client) ClientOption {
	return func(c *RESTClient) {
		c.httpClient = client
	}
}

func WithSession(ticket, csrfPreventionToken string) ClientOption {
	return func(c *RESTClient) {
		c.session = &api.Session{
			Ticket:              ticket,
			CSRFPreventionToken: csrfPreventionToken,
		}
	}
}

func WithUserPassword(username, password string) ClientOption {
	return func(c *RESTClient) {
		c.credentials = &TicketRequest{
			Username: username,
			Password: password,
		}
	}
}

func WithAPIToken(tokenid, secret string) ClientOption {
	return func(c *RESTClient) {
		c.tokenid = tokenid
		c.token = fmt.Sprintf("%s=%s", tokenid, secret)
	}
}

func (c *RESTClient) Do(ctx context.Context, httpMethod, urlPath string, req, v interface{}) error {
	endpoint := c.endpoint + urlPath

	var body io.Reader
	if req != nil {
		jsonReq, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = bytes.NewReader(jsonReq)
	}

	httpReq, err := http.NewRequestWithContext(ctx, httpMethod, endpoint, body)
	if err != nil {
		return err
	}
	httpReq.Header = c.makeAuthHeaders()
	httpReq.Header.Add("Content-Type", "application/json")

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpRsp.Body.Close()

	if err := checkResponse(httpRsp); err != nil {
		if IsNotAuthorized(err) {
			// try to remove expired ticket
			c.session = nil
		}
		return err
	}

	buf, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return err
	}

	// try unmarshalon {"data": any} firstly
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

func (c *RESTClient) makeAuthHeaders() http.Header {
	header := make(http.Header)
	header.Add("Accept", "application/json")
	if c.token != "" {
		header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s", c.token))
		header.Add("User-Agent", fmt.Sprintf("%s:%s", defaultUserAgent, c.tokenid))
	} else if c.session != nil {
		header.Add("Cookie", fmt.Sprintf("PVEAuthCookie=%s", c.session.Ticket))
		header.Add("CSRFPreventionToken", c.session.CSRFPreventionToken)
		header.Add("User-Agent", fmt.Sprintf("%s:%s", defaultUserAgent, c.session.Username))
	}
	return header
}

func (c *RESTClient) makeNewSession(ctx context.Context) error {
	var err error
	c.session, err = c.PostTicket(ctx, *c.credentials)
	if err != nil {
		return err
	}
	return nil
}

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read body while handling http response of status %d : %v", res.StatusCode, err)
	}

	if res.StatusCode == http.StatusInternalServerError || res.StatusCode == http.StatusNotImplemented {
		return &Error{code: res.StatusCode, returnMessage: res.Status}
	}
	if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
		return &Error{code: res.StatusCode, returnMessage: NotAuthorized}
	}

	if res.StatusCode == http.StatusBadRequest {
		var errorskey map[string]json.RawMessage
		if err := json.Unmarshal(body, &errorskey); err != nil {
			return err
		}
		if body, ok := errorskey["errors"]; ok {
			return fmt.Errorf("bad request: %s - %s", res.Status, body)
		}
		return fmt.Errorf("bad request: %s - %s", res.Status, string(body))
	}

	return fmt.Errorf("code: %d", res.StatusCode)
}
