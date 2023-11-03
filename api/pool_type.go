package api

type ResourcePool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment"`
}

type ResourcePoolConfig struct {
	Comment string                    `json:"comment"`
	Members []*map[string]interface{} `json:"members"`
}

type UpdateResourcePoolOption struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment,omitempty"`

	// set true for removing object from resource pool
	Delete bool `json:"delete,omitempty"`

	// array of storage name
	Storage []string `json:"storage,omitempty"`

	// array of vmid
	VMs []string `json:"vms,omitempty"`
}
