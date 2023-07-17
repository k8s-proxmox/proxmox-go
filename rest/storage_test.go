package rest

import (
	"time"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) TestGetStorages() {
	storages, err := s.restclient.GetStorages()
	if err != nil {
		s.T().Fatalf("failed to get storages: %v", err)
	}
	s.T().Logf("get storages: %v", storages)
}

func (s *TestSuite) GetTestStorage() *api.Storage {
	storages, err := s.restclient.GetStorages()
	if err != nil {
		s.T().Fatalf("failed to get storages: %v", err)
	}
	return storages[0]
}

func (s *TestSuite) TestGetStorage() {
	testStorageName := s.GetTestStorage().Storage

	storage, err := s.restclient.GetStorage(testStorageName)
	if err != nil {
		s.T().Fatalf("failed to get storage(name=%s): %v", testStorageName, err)
	}
	s.T().Logf("get storage: %v", *storage)
}

func (s *TestSuite) EnsureNoStorage(name string) {
	storage, err := s.restclient.GetStorage(name)
	if err == nil {
		s.T().Logf("error: %v", err)
		if err := s.restclient.DeleteStorage(storage.Storage); err != nil {
			s.T().Fatalf("failed to ensure no storage (name=%s): %v", storage.Storage, err)
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil && !IsNotFound(err) {
		s.T().Logf("failed to get storage(name=%s): %v", name, err)
	}
}

func (s *TestSuite) TestCreateDeleteStorage() {
	testStorageName := "test-proxmox-go"
	s.EnsureNoStorage(testStorageName)

	// create
	testOptions := api.StorageCreateOptions{
		Content: "images",
		Mkdir:   true,
		Path:    "/var/lib/vz/test",
	}
	storage, err := s.restclient.CreateStorage(testStorageName, "dir", testOptions)
	if err != nil {
		s.T().Fatalf("failed to create storage(name=%s): %v", testStorageName, err)
	}
	s.T().Logf("create storage: %v", *storage)
	time.Sleep(2 * time.Second)

	// delete
	err = s.restclient.DeleteStorage(testStorageName)
	if err != nil {
		s.T().Fatalf("failed to delete storage(name=%s): %v", testStorageName, err)
	}
}
