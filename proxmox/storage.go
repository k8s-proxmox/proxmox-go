package proxmox

type Storage struct {
	Active       int
	Avail        int
	Content      string
	Enabled      int
	Shared       int
	Storage      string
	Total        int
	Type         string
	Used         int
	UsedFraction float64 `json:"used_fraction"`
}

func (c *RESTClient) GetStorages() ([]*Storage, error) {
	var storages []*Storage
	if err := c.Get("/storage", &storages); err != nil {
		return nil, err
	}
	return storages, nil
}

func (c *RESTClient) GetStorage(name string) (*Storage, error) {
	storages, err := c.GetStorages()
	if err != nil {
		return nil, err
	}
	for _, s := range storages {
		if s.Storage == name {
			return s, nil
		}
	}
	return nil, NotFoundErr
}

// wip
// https://pve.proxmox.com/pve-docs/api-viewer/#/storage
type StorageCreateOptions struct {
	Storage     string `json:"storage,omitempty"`
	StorageType string `json:"type,omitempty"`
	// allowed cotent types
	// NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs
	Content     string `json:"content,omitempty"`
	ContentDirs string `json:"content-dirs,omitempty"`
	Format      string `json:"format,omitempty"`
	Mkdir       bool   `json:"mkdir,omitempty"`
	Path        string `json:"path,omitempty"`
}

func (c *RESTClient) CreateStorage(name, storageType string, options StorageCreateOptions) (*Storage, error) {
	options.Storage = name
	options.StorageType = storageType
	var storage *Storage
	if err := c.Post("/storage", options, &storage); err != nil {
		return nil, err
	}
	return storage, nil
}
