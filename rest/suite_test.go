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
	tokeid := os.Getenv("PROXMOX_TOKENID")
	secret := os.Getenv("PROXMOX_SECRET")
	if url == "" {
		s.T().Fatal("url must not be empty")
	}

	var loginOption ClientOption
	if user != "" && password != "" {
		loginOption = WithUserPassword(user, password)
	} else if tokeid != "" && secret != "" {
		loginOption = WithAPIToken(tokeid, secret)
	} else {
		s.T().Logf("username=%s, password=%s, tokenid=%s, secret=%s", user, password, tokeid, secret)
		s.T().Fatal("username&password or tokeid&secret pair must be provided")
	}

	base := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	restclient, err := NewRESTClient(url, loginOption, WithTransport(base))
	if err != nil {
		s.T().Logf("username=%s, password=%s, tokenid=%s, secret=%s", user, password, tokeid, secret)
		s.T().Fatalf("failed to create rest client: %v", err)
	}

	s.restclient = *restclient
}

func TestSuiteIntegration(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
