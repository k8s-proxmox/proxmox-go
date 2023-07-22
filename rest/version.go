package rest

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetVersion(ctx context.Context) (*api.Version, error) {
	var version *api.Version
	if err := c.Get(ctx, "/version", &version); err != nil {
		return nil, err
	}
	return version, nil
}
