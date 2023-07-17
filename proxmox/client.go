package proxmox

import (
	"github.com/sp-yduck/proxmox-go/api"
	"github.com/sp-yduck/proxmox-go/rest"
)

type Service struct {
	restclient *rest.RESTClient
}

func (s *Service) Nodes() ([]*api.Node, error) {
	return s.restclient.GetNodes()
}
