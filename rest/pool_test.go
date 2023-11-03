package rest

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) TestGetResourcePools() {
	pools, err := s.restclient.GetResourcePools(context.Background())
	if err != nil {
		s.T().Fatalf("failed to get pools: %v", err)
	}
	s.T().Logf("get pools: %v", pools)
}

func (s *TestSuite) TestCreateResourcePool() {
	pool := api.ResourcePool{
		PoolID:  "proxmox-go-test",
		Comment: "proxmox-go-test comment",
	}
	if err := s.restclient.CreateResourcePool(context.Background(), pool); err != nil {
		s.T().Fatalf("failed to create pool: %v", err)
	}
}

func (s *TestSuite) TestDeleteResourcePool() {
	if err := s.restclient.DeleteResourcePool(context.Background(), "proxmox-go-test"); err != nil {
		s.T().Fatalf("failed to delete pool: %v", err)
	}
}

func (s *TestSuite) TestUpdateResourcePool() {
	opts := api.UpdateResourcePoolOption{
		PoolID: "proxmox-go-test",
		VMs:    []string{"102", "103"},
		Delete: true,
	}
	if err := s.restclient.UpdateResourcePool(context.Background(), opts); err != nil {
		s.T().Fatalf("failed to update pool: %v", err)
	}
}

func (s *TestSuite) TestGetResourcePoolConfig() {
	config, err := s.restclient.GetResourcePoolConfig(context.Background(), "proxmox-go-test")
	if err != nil {
		s.T().Fatalf("failed to get pool config: %v", err)
	}
	s.T().Logf("get pool config: %v", *config)
	// for _, m := range config.Members {
	// 	s.T().Logf("get members: %v", m)
	// }
}
