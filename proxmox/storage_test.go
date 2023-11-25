package proxmox

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) TestDownloadFromURL() {
	node := "assam"
	storageName := "local"
	storage, err := s.service.Storage(context.Background(), storageName)
	if err != nil {
		s.T().Fatalf("failed to get storage: %v", err)
	}
	storage.Node = node

	opts := api.ContentDownloadOption{
		Checksum:          "f2c748fd426f4055a0c3a6d01f0282fa75aa89e514d165845fec117cb479d840",
		ChecksumAlgorithm: "sha256",
		Content:           "iso",
		Filename:          "test.img",
		URL:               "https://cloud-images.ubuntu.com/jammy/20231027/jammy-server-cloudimg-amd64-disk-kvm.img",
	}
	if err := storage.DownloadFromURL(context.Background(), opts); err != nil {
		s.T().Fatalf("failed to download content: %v", err)
	}
}
