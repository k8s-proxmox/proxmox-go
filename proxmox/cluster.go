package proxmox

import (
	"context"

	"github.com/k8s-proxmox/proxmox-go/api"
)

func (s *Service) NextID(ctx context.Context) (int, error) {
	return s.restclient.GetNextID(ctx)
}

func (s *Service) JoinConfig(ctx context.Context) (*api.ClusterJoinConfig, error) {
	return s.restclient.GetJoinConfig(ctx)
}
