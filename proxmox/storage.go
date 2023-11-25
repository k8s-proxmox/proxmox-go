package proxmox

import (
	"context"
	"errors"
	"fmt"

	"github.com/sp-yduck/proxmox-go/api"
	"github.com/sp-yduck/proxmox-go/rest"
)

type Storage struct {
	service    *Service
	restclient *rest.RESTClient
	Storage    *api.Storage
	Node       string
}

func (s *Service) Storage(ctx context.Context, name string) (*Storage, error) {
	storage, err := s.restclient.GetStorage(ctx, name)
	if err != nil {
		return nil, err
	}
	return &Storage{service: s, restclient: s.restclient, Storage: storage}, nil
}

func (s *Service) CreateStorage(ctx context.Context, name, storageType string, options api.StorageCreateOptions) (*Storage, error) {
	var storage *api.Storage
	options.Storage = name
	options.StorageType = storageType
	if err := s.restclient.Post(ctx, "/storage", options, &storage); err != nil {
		return nil, err
	}
	return &Storage{service: s, restclient: s.restclient, Storage: storage}, nil
}

func (s *Storage) Delete(ctx context.Context) error {
	return s.restclient.DeleteStorage(ctx, s.Storage.Storage)
}

func (s *Storage) GetContents(ctx context.Context) ([]*api.StorageContent, error) {
	var contents []*api.StorageContent
	if s.Node == "" {
		return nil, errors.New("Node must not be empty")
	}
	path := fmt.Sprintf("/nodes/%s/storage/%s/content", s.Node, s.Storage.Storage)
	if err := s.restclient.Get(ctx, path, &contents); err != nil {
		return nil, err
	}
	return contents, nil
}

func (s *Storage) GetContent(ctx context.Context, volumeID string) (*api.StorageContent, error) {
	contents, err := s.GetContents(ctx)
	if err != nil {
		return nil, err
	}
	for _, content := range contents {
		if content.VolID == volumeID {
			return content, nil
		}
	}
	return nil, rest.NotFoundErr
}

func (s *Storage) GetVolume(ctx context.Context, volumeID string) (*api.StorageVolume, error) {
	path := fmt.Sprintf("/nodes/%s/storage/%s/content/%s", s.Node, s.Storage.Storage, volumeID)
	var volume *api.StorageVolume
	if err := s.restclient.Get(ctx, path, &volume); err != nil {
		return nil, err
	}
	return volume, nil
}

// to do : taskid
func (s *Storage) DeleteVolume(ctx context.Context, volumeID string) error {
	path := fmt.Sprintf("/nodes/%s/storage/%s/content/%s", s.Node, s.Storage.Storage, volumeID)
	var taskid string
	if err := s.restclient.Delete(ctx, path, nil, &taskid); err != nil {
		return err
	}
	return nil
}

func (s *Storage) DownloadFromURL(ctx context.Context, opts api.ContentDownloadOption) error {
	taskid, err := s.restclient.DownloadFromURL(ctx, s.Node, s.Storage.Storage, opts)
	if err != nil {
		return err
	}
	return s.service.EnsureTaskDone(ctx, s.Node, *taskid)
}
