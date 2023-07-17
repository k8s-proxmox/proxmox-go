package rest

import (
	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetStorages() ([]*api.Storage, error) {
	var storages []*api.Storage
	if err := c.Get("/storage", &storages); err != nil {
		return nil, err
	}
	return storages, nil
}

func (c *RESTClient) GetStorage(name string) (*api.Storage, error) {
	storages, err := c.GetStorages()
	if err != nil {
		return nil, err
	}
	for _, s := range storages {
		if s.Storage == name {
			return s, nil
		}
	}
	return nil, NotFoundErr
}

func (c *RESTClient) CreateStorage(name, storageType string, options api.StorageCreateOptions) (*api.Storage, error) {
	options.Storage = name
	options.StorageType = storageType
	var storage *api.Storage
	if err := c.Post("/storage", options, &storage); err != nil {
		return nil, err
	}
	return storage, nil
}
