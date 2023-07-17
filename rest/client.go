package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/sp-yduck/proxmox-go/api"
)

type RESTClient struct {
	endpoint    string
	httpClient  *http.Client
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
		endpoint:   baseUrl,
		httpClient: &http.Client{},
	}
	for _, option := range opts {
		option(client)
	}
	if client.token == "" && client.session == nil && client.credentials != nil {
		var err error
		client.session, err = client.PostTicket(*client.credentials)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
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

func withLogin() ClientOption {
	return func(c *RESTClient) {
	}
}

func (c *RESTClient) Do(httpMethod, urlPath string, req, v interface{}) error {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url, err := url.JoinPath(c.endpoint, urlPath)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest(httpMethod, url, bytes.NewReader(jsonReq))
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
		return err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
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

func (c *RESTClient) Get(path string, res interface{}) error {
	return c.Do(http.MethodGet, path, nil, res)
}

func (c *RESTClient) Post(path string, req, res interface{}) error {
	return c.Do(http.MethodPost, path, req, res)
}

func (c *RESTClient) Put(path string, req, res interface{}) error {
	return c.Do(http.MethodPut, path, req, res)
}

func (c *RESTClient) Delete(path string, req, res interface{}) error {
	return c.Do(http.MethodDelete, path, req, res)
}

func (c *RESTClient) makeAuthHeaders() http.Header {
	header := make(http.Header)
	// header.Add("User-Agent", c.userAgent)
	header.Add("Accept", "application/json")
	if c.token != "" {
		header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s", c.token))
	} else if c.session != nil {
		header.Add("Cookie", fmt.Sprintf("PVEAuthCookie=%s", c.session.Ticket))
		header.Add("CSRFPreventionToken", c.session.CSRFPreventionToken)
	}
	return header
}

func (c *RESTClient) makeNewSession() error {
	var err error
	c.session, err = c.PostTicket(*c.credentials)
	if err != nil {
		return err
	}
	return nil
}

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	if res.StatusCode == http.StatusInternalServerError || res.StatusCode == http.StatusNotImplemented {
		return &Error{code: res.StatusCode, returnMessage: NotFound}
	}
	if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
		return &Error{code: res.StatusCode, returnMessage: NotAuthorized}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Errorf("failed to read body while handling http response of status %d : %v", res.StatusCode, err)
	}

	if res.StatusCode == http.StatusBadRequest {
		var errorskey map[string]json.RawMessage
		if err := json.Unmarshal(body, &errorskey); err != nil {
			return err
		}
		if body, ok := errorskey["errors"]; ok {
			return errors.Errorf("bad request: %s - %s", res.Status, body)
		}
		return errors.Errorf("bad request: %s - %s", res.Status, string(body))
	}

	return errors.Errorf("code: %d", res.StatusCode)
}
