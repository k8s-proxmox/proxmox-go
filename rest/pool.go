package rest

import (
	"context"
	"fmt"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetResourcePools(ctx context.Context) ([]*api.ResourcePool, error) {
	var pools []*api.ResourcePool
	if err := c.Get(ctx, "/pools", &pools); err != nil {
		return nil, err
	}
	return pools, nil
}

func (c *RESTClient) GetResourcPool(ctx context.Context, id string) (*api.ResourcePool, error) {
	pools, err := c.GetResourcePools(ctx)
	if err != nil {
		return nil, err
	}
	for _, p := range pools {
		if p.PoolID == id {
			return p, nil
		}
	}
	return nil, NotFoundErr
}

func (c *RESTClient) CreateResourcePool(ctx context.Context, pool api.ResourcePool) error {
	return c.Post(ctx, "/pools", pool, nil)
}

func (c *RESTClient) DeleteResourcePool(ctx context.Context, id string) error {
	var res map[string]interface{}
	path := fmt.Sprintf("/pools/%s", id)
	return c.Delete(ctx, path, nil, &res)
}

func (c *RESTClient) UpdateResourcePool(ctx context.Context, option api.UpdateResourcePoolOption) error {
	path := fmt.Sprintf("/pools/%s", option.PoolID)
	return c.Put(ctx, path, option, nil)
}
