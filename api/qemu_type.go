package api

import (
	"encoding/json"
)

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

type Parallel struct {
	Parallel0 string `json:"parallel0,omitempty"`
	Parallel1 string `json:"parallel1,omitempty"`
	Parallel2 string `json:"parallel2,omitempty"`
}

type Sata struct {
	Sata0 string `json:"sata0,omitempty"`
	Sata1 string `json:"sata1,omitempty"`
	Sata2 string `json:"sata2,omitempty"`
	Sata3 string `json:"sata3,omitempty"`
	Sata4 string `json:"sata4,omitempty"`
	Sata5 string `json:"sata5,omitempty"`
}

// wip n = 0~30
type Scsi struct {
	Scsi0 string `json:"scsi0,omitempty"`
}

type Serial struct {
	Serial0 string `json:"serial0,omitempty"`
}

type UnUsed struct {
	UnUsed0 string `json:"unused0,omitempty"`
}

type USB struct {
	USB0 string `json:"usb0,omitempty"`
}

type VirtIO struct {
	VirtIO0 string `json:"virtio0,omitempty"`
}

type Bool int8

const (
	True  Bool = 1
	False Bool = 0
)

// wip
// reference : https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu
type VirtualMachineCreateOptions struct {
	Acpi     Bool   `json:"acpi,omitempty"`
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
	Force Bool `json:"force,omitempty"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3).
	// Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	// Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ide
	IPConfig
	// enable/disable KVM hardware virtualization
	Kvm Bool `json:"kvm,omitempty"`
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
	OnBoot Bool `json:"onboot,omitempty"`
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
	Start Bool `json:"start,omitempty"`
	// tags of the VM. only for meta information
	Tags string `json:"tags,omitempty"`
	// enable/disable template
	Template Bool   `json:"template,omitempty"`
	VGA      string `json:"vga,omitempty"`
	// vm id
	VMID int `json:"vmid,omitempty"`
}

type VirtualMachineConfig struct {
	// Enable/disable ACPI.
	ACPI Bool `json:"acpi,omitempty"`
	// List of host cores used to execute guest processes, for example: 0,5,8-11
	Affinity string `json:"affinity,omitempty"`
	// Enable/disable communication with the QEMU Guest Agent and its properties.
	Agent string `json:"agent,omitempty"`
	// Virtual processor architecture. Defaults to the host.
	Arch Arch `json:"arch,omitempty"`
	// Arbitrary arguments passed to kvm, for example:
	// args: -no-reboot -no-hpet
	// NOTE: this option is for experts only.
	Args string `json:"args,omitempty"`
	// Configure a audio device, useful in combination with QXL/Spice.
	Audio0 string `json:"audio0,omitempty"`
	// Automatic restart after crash (currently ignored).
	AutoStart Bool `json:"autostart,omitempty"`
	// Amount of target RAM for the VM in MiB. Using zero disables the ballon driver.
	Balloon int `json:"balloon,omitempty"`
	// Select BIOS implementation.
	BIOS string `json:"bios,omitempty"`
	// boot order. ";" separated. : 'order=device1;device2;device3'
	Boot string `json:"boot,omitempty"`
	// This is an alias for option -ide2
	CDRom string `json:"cdrom,omitempty"`
	// cloud-init: Specify custom files to replace the automatically generated ones at start.
	CiCustom string `json:"cicustom,omitempty"`
	// cloud-init: Password to assign the user. Using this is generally not recommended.
	// Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords.
	CiPassword string `json:"cipassword,omitempty"`
	// Specifies the cloud-init configuration format.
	// The default depends on the configured operating system type (`ostype`.
	// We use the `nocloud` format for Linux, and `configdrive2` for windows.
	CiType string `json:"citype,omitempty"`
	// cloud-init: User name to change ssh keys and password for instead of the image's configured default user.
	CiUser string `json:"ciuser,omitempty"`
	// The number of cores per socket. : 1 ~
	Cores int `json:"cores,omitempty"`
	// emulated cpu type
	Cpu string `json:"cpu,omitempty"`
	// Limit of CPU usage.
	// NOTE: If the computer has 2 CPUs, it has total of '2' CPU time. Value '0' indicates no CPU limit.
	CpuLimit int `json:"cpulimit,omitempty"`
	// CPU weight for a VM. Argument is used in the kernel fair scheduler.
	// The larger the number is, the more CPU time this VM gets.
	// Number is relative to weights of all the other running VMs.
	CpuUnits    int    `json:"cpuunits"`
	Description string `json:"description,omitempty"`
	EfiDisk0    Bool   `json:"efidisk0,omitempty"`
	Freeze      Bool   `json:"freeze,omitempty"`
	HookScript  string `json:"hookscript,omitempty"`
	// HostPci
	HotPlug   string `json:"hotplug,omitempty"`
	HugePages string `json:"hugepages,omitempty"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3).
	// Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	// Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ide `json:"-"`
	IPConfig
	IvshMem       string `json:"ivshmem,omitempty"`
	KeepHugePages Bool   `json:"keephugepages,omitempty"`
	Keyboard      string `json:"keyboard,omitempty"`
	// enable/disable KVM hardware virtualization
	Kvm       Bool   `json:"kvm,omitempty"`
	LocalTime Bool   `json:"localtime,omitempty"`
	Lock      string `json:"lock,omitempty"`
	// specifies the QEMU machine type
	Machine string `json:"machine,omitempty"`
	// amount of RAM for the VM in MiB : 16 ~
	Memory          int         `json:"memory,omitempty"`
	MigrateDowntime json.Number `json:"migrate_downtime,omitempty"`
	MigrateSpeed    int         `json:"migrate_speed,omitempty"`
	// name for VM. Only used on the configuration web interface
	Name string `json:"name,omitempty"`
	// cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	NameServer string `json:"nameserver,omitempty"`
	// network device
	Net  `json:"-"`
	Numa int8 `json:"numa,omitempty"`
	// specifies whether a VM will be started during system bootup
	OnBoot Bool `json:"onboot,omitempty"`
	// quest OS
	OSType     OSType `json:"ostype,omitempty"`
	Parallel   `json:"-"`
	Protection Bool `json:"protection,omitempty"`
	// Allow reboot. if set to '0' the VM exit on reboot
	Reboot int    `json:"reboot,omitempty"`
	RNG0   string `json:"rng0,omitempty"`
	Sata   `json:"-"`
	// use volume as scsi hard disk or cd-rom
	// use special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume
	// use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Scsi `json:"-"`
	// SCSI controller model
	ScsiHw ScsiHw `json:"scsihw,omitempty"`
	// cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	SearchDomain string `json:"searchdomain,omitempty"`
	Serial       `json:"-"`
	Shares       int    `json:"shares,omitempty"`
	SMBios1      string `json:"smbios1,omitempty"`
	SMP          int    `json:"smp,omitempty"`
	// number of sockets
	Sockets           int    `json:"sockets,omitempty"`
	SpiceEnhancements string `json:"spice_enhancements,omitempty"`
	// cloud-init setup public ssh keys (one key per line, OpenSSH format)
	SSHKeys   string `json:"sshkeys,omitempty"`
	StartDate string `json:"startdate,omitempty"`
	StartUp   Bool   `json:"startup,omitempty"`
	Tablet    Bool   `json:"tablet,omitempty"`
	// tags of the VM. only for meta information
	Tags string `json:"tags,omitempty"`
	TDF  Bool   `json:"tdf,omitempty"`
	// enable/disable template
	Template       Bool   `json:"template,omitempty"`
	TPMState0      string `json:"tpmstate,omitempty"`
	UnUsed         `json:"-"`
	VCPUs          int    `json:"vcpus,omitempty"`
	VGA            string `json:"vga,omitempty"`
	VirtIO         `json:"-"`
	VMGenID        string `json:"vmgenid,omitempty"`
	VMStateStorage string `json:"vmstatestorage,omitempty"`
	WatchDog       string `json:"watchdog,omitempty"`
}

type VirtualMachineRebootOption struct {
	TimeOut int `json:"timeout,omitempty"`
}

type VirtualMachineResumeOption struct {
	NoCheck  Bool `json:"nocheck,omitempty"`
	SkipLock Bool `json:"skiplock,omitempty"`
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
	SkipLock      Bool   `json:"skiplock,omitempty"`
	// some command save/restore state from this location
	StateURI string `json:"stateuri,omitempty"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage.
	// Providing the special value '1' will map each source storage to itself.
	TargetStoraage string `json:"targetstorage,omitempty"`
	TimeOut        int    `json:"timeout,omitempty"`
}

type VirtualMachineStopOption struct {
	KeepActive   Bool   `json:"keepActive,omitempty"`
	MigratedFrom string `json:"migratedfrom,omitempty"`
	SkipLock     Bool   `json:"skiplock,omitempty"`
	TimeOut      int    `json:"timeout,omitempty"`
}
