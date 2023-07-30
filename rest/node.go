package rest

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetNodes(ctx context.Context) ([]*api.Node, error) {
	var nodes []*api.Node
	if err := c.Get(ctx, "/nodes", &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *RESTClient) GetNode(ctx context.Context, name string) (*api.Node, error) {
	nodes, err := c.GetNodes(ctx)
	if err != nil {
		return nil, err
	}
	for _, n := range nodes {
		if n.Node == name {
			return n, nil
		}
	}
	return nil, NotFoundErr
}

func (c *RESTClient) CreateNodeTermProxy(ctx context.Context, nodeName string, option api.TermProxyOption) (*api.TermProxy, error) {
	path := fmt.Sprintf("/nodes/%s/termproxy", nodeName)
	var termProxy *api.TermProxy
	if err := c.Post(ctx, path, option, &termProxy); err != nil {
		return nil, err
	}
	return termProxy, nil
}

func (c *RESTClient) CreateNodeVNCShell(ctx context.Context, nodeName string, option api.VNCShellOption) (*api.TermProxy, error) {
	path := fmt.Sprintf("/nodes/%s/vncshell", nodeName)
	var termProxy *api.TermProxy
	if err := c.Post(ctx, path, option, &termProxy); err != nil {
		return nil, err
	}
	return termProxy, nil
}

func (c *RESTClient) GetNodeVNCWebSocket(ctx context.Context, nodeName, port, vncticket string) (*api.VNCWebSocket, error) {
	path := fmt.Sprintf("/nodes/%s/vncwebsocket?port=%s&vncticket=%s", nodeName, port, url.QueryEscape(vncticket))
	var websocket *api.VNCWebSocket
	if err := c.Get(ctx, path, &websocket); err != nil {
		return nil, err
	}
	return websocket, nil
}
