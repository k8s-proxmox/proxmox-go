package proxmox

import (
	"crypto/sha256"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/sp-yduck/proxmox-go/rest"
)

var (
	// global Service map against sessionKeys in map[sessionKey]Session
	sessionCache sync.Map

	// mutex to control access to the GetOrCreate function to avoid duplicate
	// session creations on startup
	sessionMutex sync.Mutex
)

type Service struct {
	restclient *rest.RESTClient
}

type Params struct {
	// base endpoint of proxmox rest api
	endpoint string

	// auth config
	authConfig AuthConfig

	// rest client config
	clientConfig ClientConfig
}

type ClientConfig struct {
	InsecureSkipVerify bool
}

type AuthConfig struct {
	Username string
	Password string
	TokenID  string
	Secret   string
}

func GetOrCreateService(params Params) (*Service, error) {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	key := retrieveSessionKey(params)
	if cachedSession, ok := sessionCache.Load(key); ok {
		return cachedSession.(*Service), nil
	}

	s, err := NewService(params)
	if err != nil {
		return nil, err
	}
	sessionCache.Store(key, s)
	return s, nil
}

func NewService(params Params) (*Service, error) {
	loginOption, err := makeLoginOpts(params.authConfig)
	if err != nil {
		return nil, err
	}
	clientOptions := []rest.ClientOption{loginOption}

	if params.clientConfig.InsecureSkipVerify {
		baseClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		clientOptions = append(clientOptions, rest.WithClient(baseClient))
	}

	restclient, err := rest.NewRESTClient(params.endpoint, clientOptions...)
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

func makeLoginOpts(authConfig AuthConfig) (rest.ClientOption, error) {
	if authConfig.Username != "" && authConfig.Password != "" {
		return rest.WithUserPassword(authConfig.Username, authConfig.Password), nil
	} else if authConfig.TokenID != "" && authConfig.Secret != "" {
		return rest.WithAPIToken(authConfig.TokenID, authConfig.Secret), nil
	}
	return nil, errors.New("invalid authentication config")
}

func retrieveSessionKey(params Params) string {
	var id string
	var secret []byte
	h := sha256.New()
	if params.authConfig.Username != "" && params.authConfig.Password != "" {
		id = params.authConfig.Username
		h.Write([]byte(params.authConfig.Password))
		secret = h.Sum(nil)
	} else if params.authConfig.TokenID != "" && params.authConfig.Secret != "" {
		id = params.authConfig.TokenID
		h.Write([]byte(params.authConfig.Secret))
		secret = h.Sum(nil)
	} else {
		id = params.authConfig.Username
		h.Write([]byte(params.authConfig.Password))
		secret = h.Sum(nil)
	}
	return fmt.Sprintf("%s#%s#%x", params.endpoint, id, secret)
}
