package rest

import (
	"context"
	"fmt"

	"github.com/k8s-proxmox/proxmox-go/api"
)

func (c *RESTClient) GetStorages(ctx context.Context) ([]*api.Storage, error) {
	var storages []*api.Storage
	if err := c.Get(ctx, "/storage", &storages); err != nil {
		return nil, err
	}
	return storages, nil
}

func (c *RESTClient) GetStorage(ctx context.Context, name string) (*api.Storage, error) {
	storages, err := c.GetStorages(ctx)
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

func (c *RESTClient) CreateStorage(ctx context.Context, name, storageType string, options api.StorageCreateOptions) (*api.Storage, error) {
	options.Storage = name
	options.StorageType = storageType
	var storage *api.Storage
	if err := c.Post(ctx, "/storage", options, &storage); err != nil {
		return nil, err
	}
	return storage, nil
}

func (c *RESTClient) DeleteStorage(ctx context.Context, name string) error {
	path := fmt.Sprintf("/storage/%s", name)
	if err := c.Delete(ctx, path, nil, nil); err != nil {
		return err
	}
	return nil
}

func (c *RESTClient) DownloadFromURL(ctx context.Context, node, storage string, option api.ContentDownloadOption) (*string, error) {
	path := fmt.Sprintf("/nodes/%s/storage/%s/download-url", node, storage)
	var upid *string
	if err := c.Post(ctx, path, option, &upid); err != nil {
		return nil, err
	}
	return upid, nil
}
