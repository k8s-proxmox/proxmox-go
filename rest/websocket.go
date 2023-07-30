package rest

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) DialNodeVNCWebSocket(ctx context.Context, nodeName string, vnc api.TermProxy) (*websocket.Conn, error) {
	baseUrl := strings.Replace(c.endpoint, "https://", "wss://", 1)
	baseUrl = strings.Replace(baseUrl, "http://", "wss://", 1)
	websocketUrl := fmt.Sprintf("%s/nodes/%s/vncwebsocket?port=%s&vncticket=%s", baseUrl, nodeName, vnc.Port, url.QueryEscape(vnc.Ticket))

	conn, resp, err := c.websocketDialer().DialContext(ctx, websocketUrl, c.makeAuthHeaders())
	if err != nil {
		if resp != nil {
			return nil, errors.Errorf("failed to dial websocket: %v : %v", checkResponse(resp), err)
		}
		return nil, errors.Errorf("failed to dial websocket: %v", err)
	}

	if err := conn.WriteMessage(websocket.BinaryMessage, []byte(fmt.Sprintf("%s:%s\n", vnc.User, vnc.Ticket))); err != nil {
		return nil, errors.Errorf("failed to start session: %v", err)
	}

	return conn, nil
}

func (c *RESTClient) websocketDialer() *websocket.Dialer {
	var tlsConfig *tls.Config
	transport := c.httpClient.Transport.(*http.Transport)
	if transport != nil {
		tlsConfig = transport.TLSClientConfig
	}
	return &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 30 * time.Second,
		TLSClientConfig:  tlsConfig,
	}
}
