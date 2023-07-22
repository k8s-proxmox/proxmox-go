package rest

import (
	"context"

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
