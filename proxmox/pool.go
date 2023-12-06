package proxmox

import (
	"context"
	"fmt"
	"strings"

	"github.com/k8s-proxmox/proxmox-go/api"
)

type Pool struct {
	service *Service
	Pool    *api.ResourcePool
}

func (s *Service) Pools(ctx context.Context) ([]*api.ResourcePool, error) {
	return s.restclient.GetResourcePools(ctx)
}

func (s *Service) Pool(ctx context.Context, id string) (*Pool, error) {
	pool, err := s.restclient.GetResourcPool(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Pool{service: s, Pool: pool}, nil
}

func (s *Service) CreatePool(ctx context.Context, pool api.ResourcePool) (*Pool, error) {
	if err := s.restclient.CreateResourcePool(ctx, pool); err != nil {
		return nil, err
	}
	return s.Pool(ctx, pool.PoolID)
}

func (s *Service) DeletePool(ctx context.Context, id string) error {
	return s.restclient.DeleteResourcePool(ctx, id)
}

func IsAlreadyAPoolMemberErr(err error) bool {
	return strings.Contains(err.Error(), "is already a pool member")
}

func IsNotAPoolMemberErr(err error) bool {
	return strings.Contains(err.Error(), "is not a pool member")
}

func (p *Pool) AddVMs(ctx context.Context, vmids []int) error {
	opts := api.UpdateResourcePoolOption{
		PoolID: p.Pool.PoolID,
		VMs:    itoaSlice(vmids),
	}
	if err := p.service.restclient.UpdateResourcePool(ctx, opts); err != nil && !IsAlreadyAPoolMemberErr(err) {
		return err
	}
	return nil
}

func (p *Pool) RemoveVMs(ctx context.Context, vmids []int) error {
	opts := api.UpdateResourcePoolOption{
		PoolID: p.Pool.PoolID,
		VMs:    itoaSlice(vmids),
		Delete: true,
	}
	if err := p.service.restclient.UpdateResourcePool(ctx, opts); err != nil && !IsNotAPoolMemberErr(err) {
		return err
	}
	return nil
}

func (p *Pool) AddStorages(ctx context.Context, storageNames []string) error {
	opts := api.UpdateResourcePoolOption{
		PoolID:  p.Pool.PoolID,
		Storage: storageNames,
	}
	if err := p.service.restclient.UpdateResourcePool(ctx, opts); err != nil && !IsAlreadyAPoolMemberErr(err) {
		return err
	}
	return nil
}

func (p *Pool) RemoveStorages(ctx context.Context, storageNames []string) error {
	opts := api.UpdateResourcePoolOption{
		PoolID:  p.Pool.PoolID,
		Storage: storageNames,
		Delete:  true,
	}
	if err := p.service.restclient.UpdateResourcePool(ctx, opts); err != nil && !IsNotAPoolMemberErr(err) {
		return err
	}
	return nil
}

func (p *Pool) GetMembers(ctx context.Context) ([]*map[string]interface{}, error) {
	config, err := p.service.restclient.GetResourcePoolConfig(ctx, p.Pool.PoolID)
	if err != nil {
		return nil, err
	}
	return config.Members, nil
}

func itoaSlice(i []int) []string {
	a := []string{}
	for _, x := range i {
		a = append(a, fmt.Sprintf("%d", x))
	}
	return a
}
