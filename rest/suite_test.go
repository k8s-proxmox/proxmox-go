package rest

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	restclient RESTClient
}

func (s *TestSuite) SetupSuite() {
	url := os.Getenv("PROXMOX_URL")
	user := os.Getenv("PROXMOX_USERNAME")
	password := os.Getenv("PROXMOX_PASSWORD")
	if url == "" || user == "" || password == "" {
		s.T().Fatalf("following env var must not be empty: PROXMOX_URL=%s, POXMOX_USERNAME=%s, PROXOMOX_PASSWORD=%s", url, user, password)
	}

	base := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	restclient, err := NewRESTClient(url, WithUserPassword(user, password), WithClient(&base))
	if err != nil {
		s.T().Logf("url=%s", url)
		s.T().Logf("user=%s", user)
		s.T().Logf("password=%s", password)
		s.T().Fatalf("failed to create rest client: %v", err)
	}

	s.restclient = *restclient
}

func TestSuiteIntegration(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
