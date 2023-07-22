package proxmox

import (
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/sp-yduck/proxmox-go/rest"
)

type Service struct {
	restclient *rest.RESTClient
}

type AuthConfig struct {
	Username string
	Password string
	TokenID  string
	Secret   string
}

func NewService(url string, authConfig AuthConfig, insecure bool) (*Service, error) {
	var loginOption rest.ClientOption
	if authConfig.Username != "" && authConfig.Password != "" {
		loginOption = rest.WithUserPassword(authConfig.Username, authConfig.Password)
	} else if authConfig.TokenID != "" && authConfig.Secret != "" {
		loginOption = rest.WithAPIToken(authConfig.TokenID, authConfig.Secret)
	} else {
		return nil, errors.New("invalid authentication config")
	}
	clientOptions := []rest.ClientOption{loginOption}

	if insecure {
		baseClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		clientOptions = append(clientOptions, rest.WithClient(baseClient))
	}

	restclient, err := rest.NewRESTClient(url, clientOptions...)
	if err != nil {
		return nil, err
	}
	return &Service{restclient: restclient}, nil
}

func NewServiceWithUserPassword(url, user, password string, insecure bool) (*Service, error) {
	clientOptions := []rest.ClientOption{
		rest.WithUserPassword(user, password),
	}

	if insecure {
		baseClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		clientOptions = append(clientOptions, rest.WithClient(baseClient))
	}

	restclient, err := rest.NewRESTClient(url, clientOptions...)
	if err != nil {
		return nil, err
	}
	return &Service{restclient: restclient}, nil
}

func NewServiceWithAPIToken(url, tokenid, secret string, insecure bool) (*Service, error) {
	clientOptions := []rest.ClientOption{
		rest.WithAPIToken(tokenid, secret),
	}
	if insecure {
		baseClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		clientOptions = append(clientOptions, rest.WithClient(baseClient))
	}

	restclient, err := rest.NewRESTClient(url, clientOptions...)
	if err != nil {
		return nil, err
	}
	return &Service{restclient: restclient}, nil
}

func (s *Service) RESTClient() *rest.RESTClient {
	return s.restclient
}
