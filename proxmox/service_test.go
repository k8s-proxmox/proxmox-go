package proxmox

import (
	"os"
	"testing"
)

func TestGetOrCreateService(t *testing.T) {
	url := os.Getenv("PROXMOX_URL")
	user := os.Getenv("PROXMOX_USERNAME")
	password := os.Getenv("PROXMOX_PASSWORD")
	tokeid := os.Getenv("PROXMOX_TOKENID")
	secret := os.Getenv("PROXMOX_SECRET")
	method := os.Getenv("PROXMOX_AUTH_METHOD")
	if url == "" {
		t.Fatal("url must not be empty")
	}

	params := Params{
		endpoint: url,
		authConfig: AuthConfig{
			AuthMethod: method,
			Username:   user,
			Password:   password,
			TokenID:    tokeid,
			Secret:     secret,
		},
		clientConfig: ClientConfig{
			InsecureSkipVerify: true,
		},
	}

	var svc *Service
	for i := 0; i < 10; i++ {
		s, err := GetOrCreateService(params)
		if err != nil {
			t.Fatalf("failed to get/create service: %v", err)
		}
		if i > 0 && s != svc {
			t.Fatalf("should not create new service: %v(cached)!=%v(new)", svc, s)
		}
		svc = s
	}
}
