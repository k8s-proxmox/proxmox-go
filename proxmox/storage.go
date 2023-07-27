package proxmox

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
	"github.com/sp-yduck/proxmox-go/rest"
)

type Storage struct {
	restclient *rest.RESTClient
	Storage    *api.Storage
}

func (s *Service) Storage(ctx context.Context, name string) (*Storage, error) {
	storage, err := s.restclient.GetStorage(ctx, name)
	if err != nil {
		return nil, err
	}
	return &Storage{restclient: s.restclient, Storage: storage}, nil
}

func (s *Service) CreateStorage(ctx context.Context, name, storageType string, options api.StorageCreateOptions) (*Storage, error) {
	var storage *api.Storage
	options.Storage = name
	options.StorageType = storageType
	if err := s.restclient.Post(ctx, "/storage", options, &storage); err != nil {
		return nil, err
	}
	return &Storage{restclient: s.restclient, Storage: storage}, nil
}
