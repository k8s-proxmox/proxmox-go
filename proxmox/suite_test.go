package proxmox

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	service Service
}

func (s *TestSuite) SetupSuite() {
	url := os.Getenv("PROXMOX_URL")
	user := os.Getenv("PROXMOX_USERNAME")
	password := os.Getenv("PROXMOX_PASSWORD")
	tokeid := os.Getenv("PROXMOX_TOKENID")
	secret := os.Getenv("PROXMOX_SECRET")
	if url == "" {
		s.T().Fatal("url must not be empty")
	}

	authConfig := AuthConfig{
		Username: user,
		Password: password,
		TokenID:  tokeid,
		Secret:   secret,
	}

	service, err := NewService(url, authConfig, true)
	if err != nil {
		s.T().Fatalf("failed to create new service: %v", err)
	}
	s.service = *service
}

func TestSuiteIntegration(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
