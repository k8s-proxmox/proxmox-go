package proxmox

type Node struct {
	Cpu            float32 `json:"cpu"`
	Disk           int     `json:"disk"`
	ID             string  `json:"id"`
	Level          string  `json:"level"`
	MaxCpu         int     `json:"maxcpu"`
	MaxDisk        int     `json:"maxdisk"`
	MaxMem         int     `json:"maxmem"`
	Mem            int     `json:"mem"`
	Node           string  `json:"node"`
	SSLFingerprint string  `json:"ssl_fingerprint"`
	Stauts         string  `json:"status"`
	Type           string  `json:"type"`
	UpTime         int     `json:"uptime"`
}

func (c *RESTClient) GetNodes() ([]*Node, error) {
	var nodes []*Node
	if err := c.Get("/nodes", &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *RESTClient) GetNode(name string) (*Node, error) {
	nodes, err := c.GetNodes()
	if err != nil {
		return nil, err
	}
	for _, n := range nodes {
		if n.Node == name {
			return n, nil
		}
	}
	return nil, NotFoundErr
}
