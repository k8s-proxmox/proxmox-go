package api

import (
	"encoding/json"
)

type Storage struct {
	Active       int     `json:"active"`
	Avail        int     `json:"avail"`
	Content      string  `json:"content"`
	Enabled      int     `json:"enabled"`
	Shared       int     `json:"shared"`
	Storage      string  `json:"storage"`
	Total        int     `json:"total"`
	Type         string  `json:"type"`
	Used         int     `json:"used"`
	UsedFraction float64 `json:"used_fraction"`
}

// wip
// https://pve.proxmox.com/pve-docs/api-viewer/#/storage
type StorageCreateOptions struct {
	// required. storage identifier
	Storage string `json:"storage,omitempty"`

	// required. btrfs,cephfs,cifs,dir,glusterfs,iscsi,iscsidirect,lvm,lvmthin,nfs,pbs,rbd,zfs,zfspool
	StorageType string `json:"type,omitempty"`

	Authsupported string `json:"authsupported,omitempty"`

	// base volume. this volume is automatically activated
	Base string `json:"base,omitempty"`

	BlockSize string `json:"blocksize,omitempty"`

	// set I/O bandwidth limit for various operations (in KiB/s)
	BWLimit string `json:"bwlimit,omitempty"`

	// host group for comstar views
	ComstarHG string `json:"comstar_hg,omitempty"`

	// target group for comstar views
	ComstarTG string `json:"comstar_tg,omitempty"`

	// allowed content types
	// NOTE: the value 'rootdir' is used for Containers, and value 'images' for VMs
	Content string `json:"content,omitempty"`

	// overrides for default content type directories
	ContentDirs string `json:"content-dirs,omitempty"`

	// create base directory if it doesn't exist
	CreateBasePath *bool `json:"create-base-path,omitempty"`

	// populate directory with the default structure
	CreateSubDirs *bool `json:"create-subdirs,omitempty"`

	DataPool string `json:"data-pool,omitempty"`

	// proxmox backup serverr datastore name
	DataStore string `json:"datastore,omitempty"`

	// CIFS domain
	Domain string `json:"domain,omitempty"`

	// NFS export path
	Export string `json:"export,omitempty"`

	Format string `json:"format,omitempty"`

	// iscsi provider
	ISCSIProvider string `json:"iscsiprovider,omitempty"`

	Mkdir *bool `json:"mkdir,omitempty"`

	// filesystem path
	Path string `json:"path,omitempty"`

	// iSCSI portal (ip or dns name with optional port)
	Portal string `json:"portal,omitempty"`

	Pool string `json:"pool,omitempty"`

	// server ip or domain
	Server string `json:"server,omitempty"`

	// CIFS share
	Share string `json:"share,omitempty"`

	// iSCSI target
	Target string `json:"target,omitempty"`

	// lvm thin pool lv name
	ThinPool string `json:"thinpool,omitempty"`

	// volume group name
	VGName string `json:"vgname,omitempty"`

	// Glusterfs volume
	Volume string `json:"volume,omitempty"`
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
