package proxmox

type Version struct {
	Release string  `json:"release"`
	RepoID  string  `json:"repoid"`
	Version string  `json:"version"`
	Console Console `json:"console"`
}

type Console string

func (c *RESTClient) GetVersion() (*Version, error) {
	var version *Version
	if err := c.Get("/version", &version); err != nil {
		return nil, err
	}
	return version, nil
}
