package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/k8s-proxmox/proxmox-go/api"
)

// implementation of http.RoundTripper
type Transport struct {
	AuthProvider AuthProvider
	Base         http.RoundTripper
}

func (t *Transport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// set base Transport
func (t *Transport) SetBase(base http.RoundTripper) {
	t.Base = base
}

func (t *Transport) addAuthHeader(header *http.Header) error {
	return t.AuthProvider.addAuthHeader(header)
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqBodyClosed := false
	if req.Body != nil {
		defer func() {
			if !reqBodyClosed {
				req.Body.Close()
			}
		}()
	}

	if t.AuthProvider == nil {
		return nil, fmt.Errorf("proxmox-go: Transport's AuthProvider is nil")
	}

	req2 := cloneRequest(req)
	if err := t.addAuthHeader(&req2.Header); err != nil {
		return nil, fmt.Errorf("proxmox-go: Authentication faild: %v", err)
	}
	reqBodyClosed = true
	return t.base().RoundTrip(req2)
}

func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}

// AuthProvder is responsible for setting 'valid' auth header
// to http.Request.
type AuthProvider interface {
	// add auth header to provided header
	addAuthHeader(header *http.Header) error
}

// AuthProvider implementation
type TokenProvider struct {
	tokenID string
	secret  string
}

func NewTokenProvider(tokenID, secret string) *TokenProvider {
	return &TokenProvider{
		tokenID: tokenID,
		secret:  secret,
	}
}

func (t *TokenProvider) ID() string {
	return t.tokenID
}

func (t *TokenProvider) APIToken() string {
	return fmt.Sprintf("%s=%s", t.tokenID, t.secret)
}

func (t *TokenProvider) addAuthHeader(header *http.Header) error {
	if t.ID() == "" && t.secret == "" {
		return fmt.Errorf("tokenid and secret must not be empty")
	}
	header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s", t.APIToken()))
	header.Add("User-Agent", fmt.Sprintf("%s:%s", defaultUserAgent, t.ID()))
	return nil
}

// AuthProvider implementation
type TicketProvider struct {
	baseUrl       string
	baseTransport http.RoundTripper
	username      string
	password      string
	session       *api.Session
	expiry        time.Time
	mu            sync.Mutex
}

func NewTicketProvider(baseTransport http.RoundTripper, baseUrl, username, password string) *TicketProvider {
	t := &TicketProvider{
		baseUrl:       baseUrl,
		baseTransport: baseTransport,
		username:      username,
		password:      password,
		expiry:        time.Now(),
		mu:            sync.Mutex{},
	}
	return t
}

func (t *TicketProvider) base() http.RoundTripper {
	if t.baseTransport != nil {
		return t.baseTransport
	}
	return http.DefaultTransport
}

func (t *TicketProvider) ExpiryDelta() time.Duration {
	return t.expiry.Sub(time.Now())
}

// check ticket expiry and refresh it if expired
// then return valid session
func (t *TicketProvider) Session() (*api.Session, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.ExpiryDelta() < 5*time.Minute {
		req := TicketRequest{Username: t.username, Password: t.password}
		if err := t.startNewSession(req); err != nil {
			return nil, err
		}
	}
	return t.session, nil
}

func (t *TicketProvider) startNewSession(req TicketRequest) error {
	session, err := t.retrieveSessionTokens(req)
	if err != nil {
		return err
	}
	t.session = session

	// Tickets have a limited lifetime of 2 hours.
	// https://pve.proxmox.com/wiki/Proxmox_VE_API#Ticket_Cookie
	t.expiry = time.Now().Add(2 * time.Hour)
	return nil
}

func (t *TicketProvider) retrieveSessionTokens(req TicketRequest) (*api.Session, error) {
	endpoint := t.baseUrl + "/access/ticket"
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(jsonReq)
	httpReq, err := http.NewRequestWithContext(context.Background(), http.MethodPost, endpoint, body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	client := &http.Client{Transport: t.base()}
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	buf, err := checkResponse(httpRsp)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session token: %v", err)
	}
	var datakey map[string]json.RawMessage
	if err := json.Unmarshal(buf, &datakey); err != nil {
		return nil, err
	}
	var session *api.Session
	data := datakey["data"]
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}
	return session, nil
}

func (t *TicketProvider) addAuthHeader(header *http.Header) error {
	session, err := t.Session()
	if err != nil {
		return err
	}
	header.Add("Cookie", fmt.Sprintf("PVEAuthCookie=%s", session.Ticket))
	header.Add("CSRFPreventionToken", session.CSRFPreventionToken)
	header.Add("User-Agent", fmt.Sprintf("%s:%s", defaultUserAgent, session.Username))
	return nil
}
