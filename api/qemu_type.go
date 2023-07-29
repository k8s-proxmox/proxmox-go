package api

type VirtualMachine struct {
	Cpu       float32       `json:",omitempty"`
	Cpus      int           `json:"cpus,omitempty"`
	Disk      int           `json:"disk,omitempty"`
	DiskRead  int           `json:"diskread,omitempty"`
	DiskWrite int           `json:"diskwrite,omitempty"`
	MaxDisk   int           `json:"maxdisk,omitempty"`
	MaxMem    int           `json:"maxmem,omitempty"`
	Mem       int           `json:"mem,omitempty"`
	Name      string        `json:"name,omitempty"`
	NetIn     int           `json:"netin,omitempty"`
	NetOut    int           `json:"netout,omitempty"`
	Status    ProcessStatus `json:"status,omitempty"`
	Template  int           `json:"template,omitempty"`
	UpTime    int           `json:"uptime,omitempty"`
	VMID      int           `json:"vmid,omitempty"`
}

type ProcessStatus string

const (
	ProcessStatusRunning ProcessStatus = "running"
	ProcessStatusStopped ProcessStatus = "stopped"
	ProcessStatusPaused  ProcessStatus = "paused"
)

type Arch string
type OSType string
type ScsiHw string

const (
	X86_64  Arch = "x86_64"
	Aarch64 Arch = "aarch64"
)

const (
	Other OSType = "other"
	Wxp
	W2k
	W2k3
	W2k8
	Wvista
	Win7
	Win8
	Win10
	Win11
	// linux 2.4 kernel
	L24 OSType = "l24"
	// linux 2.6-6 kernel
	L26     OSType = "l26"
	Solaris OSType = "solaris"
)

const (
	Lsi              = "lsi"
	Lsi53c810        = "lsi53c810"
	VirtioScsiPci    = "virtio-scsi-pci"
	VirtioScsiSingle = "virtio-scsi-single"
	Megasas          = "megasas"
	Pvscsi           = "pvscsi"
)

type Ide struct {
	Ide0 string `json:"ide0,omitempty"`
	Ide1 string `json:"ide1,omitempty"`
	Ide2 string `json:"ide2,omitempty"`
}

type IPConfig struct {
	IPConfig0 string `json:"ipconfig0,omitempty"`
}

type Net struct {
	Net0 string `json:"net0,omitempty"`
}

// wip n = 0~30
type Scsi struct {
	Scsi0 string `json:"scsi0,omitempty"`
}

type Serial struct {
	Serial0 string `json:"serial0,omitempty"`
}

// wip
// reference : https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu
type VirtualMachineCreateOptions struct {
	Acpi     bool   `json:"acpi,omitempty"`
	Affinity string `json:"affinity,omitempty"`
	Agent    string `json:"agent,omitempty"`
	Arch     Arch   `json:"arch,omitempty"`
	// boot order. ";" separated. : 'order=device1;device2;device3'
	Boot string `json:"boot,omitempty"`
	// cloud-init custom files
	CiCustom   string `json:"cicustom,omitempty"`
	CiPassword string `json:"cipassword,omitempty"`
	CiType     string `json:"citype,omitempty"`
	CiUser     string `json:"ciuser,omitempty"`
	// number of cores : 1 ~
	Cores int `json:"cores,omitempty"`
	// emulated cpu type
	Cpu string `json:"cpu,omitempty"`
	// limit of cpu usage : 0 ~
	// 0 indicated no limit
	CpuLimit    int    `json:"cpulimit,omitempty"`
	Description string `json:"description,omitempty"`

	// allow to overwrite existing VM
	Force bool `json:"force,omitempty"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3).
	// Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	// Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ide
	IPConfig
	// enable/disable KVM hardware virtualization
	Kvm bool `json:"kvm,omitempty"`
	// specifies the QEMU machine type
	Machine string `json:"machine,omitempty"`
	// amount of RAM for the VM in MiB : 16 ~
	Memory int `json:"memory,omitempty"`
	// name for VM. Only used on the configuration web interface
	Name string `json:"name,omitempty"`
	// cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	NameServer string `json:"nameserver,omitempty"`
	// network device
	Net
	// specifies whether a VM will be started during system bootup
	OnBoot bool `json:"onboot,omitempty"`
	// quest OS
	OSType OSType `json:"ostype,omitempty"`
	// Allow reboot. if set to '0' the VM exit on reboot
	Reboot int `json:"reboot,omitempty"`
	// use volume as scsi hard disk or cd-rom
	// use special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume
	// use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Scsi
	// SCSI controller model
	ScsiHw ScsiHw `json:"scsihw,omitempty"`
	// cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	SearchDomain string `json:"searchdomain,omitempty"`
	Serial
	// number of sockets
	Sockets int `json:"sockets,omitempty"`
	// cloud-init setup public ssh keys (one key per line, OpenSSH format)
	SSHKeys string `json:"sshkeys,omitempty"`
	// start VM after it was created successfully
	Start bool `json:"start,omitempty"`
	// tags of the VM. only for meta information
	Tags string `json:"tags,omitempty"`
	// enable/disable template
	Template bool   `json:"template,omitempty"`
	VGA      string `json:"vga,omitempty"`
	// vm id
	VMID int `json:"vmid,omitempty"`
}

type VirtualMachineConfig struct {
	// PVE Metadata
	Digest      string `json:"digest"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Meta        string `json:"meta,omitempty"`
	VMGenID     string `json:"vmgenid,omitempty"`
	Hookscript  string `json:"hookscript,omitempty"`
	Hotplug     string `json:"hotplug,omitempty"`
	Template    int    `json:"template,omitempty"`

	Tags string `json:"tags,omitempty"`

	Protection int    `json:"protection,omitempty"`
	Lock       string `json:"lock,omitempty"`

	// Boot configuration
	Boot   string `json:"boot,omitempty"`
	OnBoot int    `json:"onboot,omitempty"`

	// Qemu general specs
	OSType  string `json:"ostype,omitempty"`
	Machine string `json:"machine,omitempty"`
	Args    string `json:"args,omitempty"`

	// Qemu firmware specs
	Bios     string `json:"bios,omitempty"`
	EFIDisk0 string `json:"efidisk0,omitempty"`
	SMBios1  string `json:"smbios1,omitempty"`
	Acpi     int    `json:"acpi,omitempty"`

	// Qemu CPU specs
	Sockets  int    `json:"sockets,omitempty"`
	Cores    int    `json:"cores,omitempty"`
	CPU      string `json:"cpu,omitempty"`
	CPULimit int    `json:"cpulimit,omitempty"`
	CPUUnits int    `json:"cpuunits,omitempty"`
	Vcpus    int    `json:"vcpus,omitempty"`
	Affinity string `json:"affinity,omitempty"`

	// Qemu memory specs
	Numa      int    `json:"numa,omitempty"`
	Memory    int    `json:"memory,omitempty"`
	Hugepages string `json:"hugepages,omitempty"`
	Balloon   int    `json:"balloon,omitempty"`

	// Other Qemu devices
	VGA       string `json:"vga,omitempty"`
	SCSIHW    string `json:"scsihw,omitempty"`
	TPMState0 string `json:"tpmstate0,omitempty"`
	Rng0      string `json:"rng0,omitempty"`
	Audio0    string `json:"audio0,omitempty"`

	// Disk devices
	Ide

	Scsi

	// Sata
	// Virtio
	// Unused

	// Network devices
	Net

	// NUMA
	// Host PCI devices HostPci

	// Serial devices
	Serial

	// USB devices
	// Parallel devices
	// Cloud-init
	CIType       string `json:"citype,omitempty"`
	CIUser       string `json:"ciuser,omitempty"`
	CIPassword   string `json:"cipassword,omitempty"`
	Nameserver   string `json:"nameserver,omitempty"`
	Searchdomain string `json:"searchdomain,omitempty"`
	SSHKeys      string `json:"sshkeys,omitempty"`
	CICustom     string `json:"cicustom,omitempty"`

	// Cloud-init interfaces
	// IPConfig
}

type VirtualMachineRebootOption struct {
	TimeOut int `json:"timeout,omitempty"`
}

type VirtualMachineResumeOption struct {
	NoCheck  bool `json:"nocheck,omitempty"`
	SkipLock bool `json:"skiplock,omitempty"`
}

type VirtualMachineStartOption struct {
	// override qemu's -cpu argument with the given string
	ForceCPU string `json:"force-cpu,omitempty"`
	// specifies the qemu machine type
	Machine string `json:"machine,omitempty"`
	// cluster node name
	MigratedFroom string `json:"migratedfrom,omitempty"`
	// cidr of (sub) network that is used for migration
	MigrationNetwork string `json:"migration_network,omitempty"`
	// migration traffic is ecrypted using an SSH tunnel by default.
	// On secure, completely private networks this can be disabled to increase performance.
	MigrationType string `json:"migration_type,omitempty"`
	SkipLock      bool   `json:"skiplock,omitempty"`
	// some command save/restore state from this location
	StateURI string `json:"stateuri,omitempty"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage.
	// Providing the special value '1' will map each source storage to itself.
	TargetStoraage string `json:"targetstorage,omitempty"`
	TimeOut        int    `json:"timeout,omitempty"`
}

type VirtualMachineStopOption struct {
	KeepActive   bool   `json:"keepActive,omitempty"`
	MigratedFrom string `json:"migratedfrom,omitempty"`
	SkipLock     bool   `json:"skiplock,omitempty"`
	TimeOut      int    `json:"timeout,omitempty"`
}
