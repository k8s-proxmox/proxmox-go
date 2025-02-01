package proxmox

import (
	"context"

	"github.com/k8s-proxmox/proxmox-go/api"
	"github.com/k8s-proxmox/proxmox-go/rest"
)

type Node struct {
	service    *Service
	restclient *rest.RESTClient
	Node       *api.Node
}

func (s *Service) GetNodes(ctx context.Context) ([]*api.Node, error) {
	return s.restclient.GetNodes(ctx)
}

func (s *Service) GetNode(ctx context.Context, name string) (*api.Node, error) {
	return s.restclient.GetNode(ctx, name)
}

func (s *Service) Node(ctx context.Context, name string) (*Node, error) {
	node, err := s.restclient.GetNode(ctx, name)
	if err != nil {
		return nil, err
	}
	return &Node{service: s, restclient: s.restclient, Node: node}, nil
}

func (n *Node) GetStorages(ctx context.Context) ([]*api.Storage, error) {
	return n.restclient.GetNodeStorages(ctx, n.Node.Node)
}
