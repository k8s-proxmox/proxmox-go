package api

import (
	"encoding/json"
)

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

type StorageContent struct {
	Storage string `json:",omitempty"`
	Content string `json:",omitempty"`
	// to do : use custom type instead of json.Number
	CTime     json.Number `json:",omitempty"`
	Encrypted string
	Format    string
	Notes     string
	Parent    string
	Protected bool
	Size      int
	Used      int
	// to do : Verificateion
	VMID  int
	VolID string `josn:"volid,omitempty"`
}

type StorageVolume struct {
	Format string `json:",omitempty"`
	Path   string `json:",omitempty"`
	Size   int    `json:",omitempty"`
	Used   int    `json:",omitempty"`
}
