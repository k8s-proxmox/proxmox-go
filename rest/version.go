package rest

import (
	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetVersion() (*api.Version, error) {
	var version *api.Version
	if err := c.Get("/version", &version); err != nil {
		return nil, err
	}
	return version, nil
}
