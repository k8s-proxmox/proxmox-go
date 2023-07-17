package rest

import (
	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetNodes() ([]*api.Node, error) {
	var nodes []*api.Node
	if err := c.Get("/nodes", &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *RESTClient) GetNode(name string) (*api.Node, error) {
	nodes, err := c.GetNodes()
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
