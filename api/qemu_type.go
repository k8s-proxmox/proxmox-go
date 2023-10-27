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
	Ide3 string `json:"ide3,omitempty"`
}

type IPConfig struct {
	IPConfig0  string `json:"ipconfig0,omitempty"`
	IPConfig1  string `json:"ipconfig1,omitempty"`
	IPConfig2  string `json:"ipconfig2,omitempty"`
	IPConfig3  string `json:"ipconfig3,omitempty"`
	IPConfig4  string `json:"ipconfig4,omitempty"`
	IPConfig5  string `json:"ipconfig5,omitempty"`
	IPConfig6  string `json:"ipconfig6,omitempty"`
	IPConfig7  string `json:"ipconfig7,omitempty"`
	IPConfig8  string `json:"ipconfig8,omitempty"`
	IPConfig9  string `json:"ipconfig9,omitempty"`
	IPConfig10 string `json:"ipconfig10,omitempty"`
	IPConfig11 string `json:"ipconfig11,omitempty"`
	IPConfig12 string `json:"ipconfig12,omitempty"`
	IPConfig13 string `json:"ipconfig13,omitempty"`
	IPConfig14 string `json:"ipconfig14,omitempty"`
	IPConfig15 string `json:"ipconfig15,omitempty"`
	IPConfig16 string `json:"ipconfig16,omitempty"`
	IPConfig17 string `json:"ipconfig17,omitempty"`
	IPConfig18 string `json:"ipconfig18,omitempty"`
	IPConfig19 string `json:"ipconfig19,omitempty"`
	IPConfig20 string `json:"ipconfig20,omitempty"`
	IPConfig21 string `json:"ipconfig21,omitempty"`
	IPConfig22 string `json:"ipconfig22,omitempty"`
	IPConfig23 string `json:"ipconfig23,omitempty"`
	IPConfig24 string `json:"ipconfig24,omitempty"`
	IPConfig25 string `json:"ipconfig25,omitempty"`
	IPConfig26 string `json:"ipconfig26,omitempty"`
	IPConfig27 string `json:"ipconfig27,omitempty"`
	IPConfig28 string `json:"ipconfig28,omitempty"`
	IPConfig29 string `json:"ipconfig29,omitempty"`
	IPConfig30 string `json:"ipconfig30,omitempty"`
	IPConfig31 string `json:"ipconfig31,omitempty"`
}

type Net struct {
	Net0  string `json:"net0,omitempty"`
	Net1  string `json:"net1,omitempty"`
	Net2  string `json:"net2,omitempty"`
	Net3  string `json:"net3,omitempty"`
	Net4  string `json:"net4,omitempty"`
	Net5  string `json:"net5,omitempty"`
	Net6  string `json:"net6,omitempty"`
	Net7  string `json:"net7,omitempty"`
	Net8  string `json:"net8,omitempty"`
	Net9  string `json:"net9,omitempty"`
	Net10 string `json:"net10,omitempty"`
	Net11 string `json:"net11,omitempty"`
	Net12 string `json:"net12,omitempty"`
	Net13 string `json:"net13,omitempty"`
	Net14 string `json:"net14,omitempty"`
	Net15 string `json:"net15,omitempty"`
	Net16 string `json:"net16,omitempty"`
	Net17 string `json:"net17,omitempty"`
	Net18 string `json:"net18,omitempty"`
	Net19 string `json:"net19,omitempty"`
	Net20 string `json:"net20,omitempty"`
	Net21 string `json:"net21,omitempty"`
	Net22 string `json:"net22,omitempty"`
	Net23 string `json:"net23,omitempty"`
	Net24 string `json:"net24,omitempty"`
	Net25 string `json:"net25,omitempty"`
	Net26 string `json:"net26,omitempty"`
	Net27 string `json:"net27,omitempty"`
	Net28 string `json:"net28,omitempty"`
	Net29 string `json:"net29,omitempty"`
	Net30 string `json:"net30,omitempty"`
	Net31 string `json:"net31,omitempty"`
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

type Scsi struct {
	Scsi0  string `json:"scsi0,omitempty"`
	Scsi1  string `json:"scsi1,omitempty"`
	Scsi2  string `json:"scsi2,omitempty"`
	Scsi3  string `json:"scsi3,omitempty"`
	Scsi4  string `json:"scsi4,omitempty"`
	Scsi5  string `json:"scsi5,omitempty"`
	Scsi6  string `json:"scsi6,omitempty"`
	Scsi7  string `json:"scsi7,omitempty"`
	Scsi8  string `json:"scsi8,omitempty"`
	Scsi9  string `json:"scsi9,omitempty"`
	Scsi10 string `json:"scsi10,omitempty"`
	Scsi11 string `json:"scsi11,omitempty"`
	Scsi12 string `json:"scsi12,omitempty"`
	Scsi13 string `json:"scsi13,omitempty"`
	Scsi14 string `json:"scsi14,omitempty"`
	Scsi15 string `json:"scsi15,omitempty"`
	Scsi16 string `json:"scsi16,omitempty"`
	Scsi17 string `json:"scsi17,omitempty"`
	Scsi18 string `json:"scsi18,omitempty"`
	Scsi19 string `json:"scsi19,omitempty"`
	Scsi20 string `json:"scsi20,omitempty"`
	Scsi21 string `json:"scsi21,omitempty"`
	Scsi22 string `json:"scsi22,omitempty"`
	Scsi23 string `json:"scsi23,omitempty"`
	Scsi24 string `json:"scsi24,omitempty"`
	Scsi25 string `json:"scsi25,omitempty"`
	Scsi26 string `json:"scsi26,omitempty"`
	Scsi27 string `json:"scsi27,omitempty"`
	Scsi28 string `json:"scsi28,omitempty"`
	Scsi29 string `json:"scsi29,omitempty"`
	Scsi30 string `json:"scsi30,omitempty"`
}

type Serial struct {
	Serial0 string `json:"serial0,omitempty"`
	Serial1 string `json:"serial1,omitempty"`
	Serial2 string `json:"serial2,omitempty"`
	Serial3 string `json:"serial3,omitempty"`
}

type UnUsed struct {
	UnUsed0 string `json:"unused0,omitempty"`
	UnUsed1 string `json:"unused1,omitempty"`
	UnUsed2 string `json:"unused2,omitempty"`
	UnUsed3 string `json:"unused3,omitempty"`
	UnUsed4 string `json:"unused4,omitempty"`
	UnUsed5 string `json:"unused5,omitempty"`
	UnUsed6 string `json:"unused6,omitempty"`
	UnUsed7 string `json:"unused7,omitempty"`
}

type USB struct {
	USB0  string `json:"usb0,omitempty"`
	USB1  string `json:"usb1,omitempty"`
	USB2  string `json:"usb2,omitempty"`
	USB3  string `json:"usb3,omitempty"`
	USB4  string `json:"usb4,omitempty"`
	USB5  string `json:"usb5,omitempty"`
	USB6  string `json:"usb6,omitempty"`
	USB7  string `json:"usb7,omitempty"`
	USB8  string `json:"usb8,omitempty"`
	USB9  string `json:"usb9,omitempty"`
	USB10 string `json:"usb10,omitempty"`
	USB11 string `json:"usb11,omitempty"`
	USB12 string `json:"usb12,omitempty"`
	USB13 string `json:"usb13,omitempty"`
	USB14 string `json:"usb14,omitempty"`
}

type VirtIO struct {
	VirtIO0  string `json:"virtio0,omitempty"`
	VirtIO1  string `json:"virtio1,omitempty"`
	VirtIO2  string `json:"virtio2,omitempty"`
	VirtIO3  string `json:"virtio3,omitempty"`
	VirtIO4  string `json:"virtio4,omitempty"`
	VirtIO5  string `json:"virtio5,omitempty"`
	VirtIO6  string `json:"virtio6,omitempty"`
	VirtIO7  string `json:"virtio7,omitempty"`
	VirtIO8  string `json:"virtio8,omitempty"`
	VirtIO9  string `json:"virtio9,omitempty"`
	VirtIO10 string `json:"virtio10,omitempty"`
	VirtIO11 string `json:"virtio11,omitempty"`
	VirtIO12 string `json:"virtio12,omitempty"`
	VirtIO13 string `json:"virtio13,omitempty"`
	VirtIO14 string `json:"virtio14,omitempty"`
	VirtIO15 string `json:"virtio15,omitempty"`
}

type HostPci struct {
	HostPci0 string `json:"hostpci0,omitempty"`
	HostPci1 string `json:"hostpci1,omitempty"`
	HostPci2 string `json:"hostpci2,omitempty"`
	HostPci3 string `json:"hostpci3,omitempty"`
}

type NumaS struct {
	Numa0 string `json:"numa0,omitempty"`
	Numa1 string `json:"numa1,omitempty"`
	Numa2 string `json:"numa2,omitempty"`
	Numa3 string `json:"numa3,omitempty"`
	Numa4 string `json:"numa4,omitempty"`
	Numa5 string `json:"numa5,omitempty"`
	Numa6 string `json:"numa6,omitempty"`
	Numa7 string `json:"numa7,omitempty"`
}

// reference : https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/qemu
type VirtualMachineCreateOptions struct {
	// Enable/disable ACPI.
	ACPI int8 `json:"acpi,omitempty"`
	// List of host cores used to execute guest processes, for example: 0,5,8-11
	Affinity string `json:"affinity,omitempty"`
	// Enable/disable communication with the QEMU Guest Agent and its properties.
	Agent string `json:"agent,omitempty"`
	// Virtual processor architecture. Defaults to the host.
	Arch Arch `json:"arch,omitempty"`
	// The backup archive. Either the file system path to a .tar or .vma file
	// (use '-' to pipe data from stdin) or a proxmox storage backup volume identifier.
	Archive string `json:"archive,omitempty"`
	// Arbitrary arguments passed to kvm, for example:
	// args: -no-reboot -no-hpet
	// NOTE: this option is for experts only.
	Args string `json:"args,omitempty"`
	// Configure a audio device, useful in combination with QXL/Spice.
	Audio0 string `json:"audio0,omitempty"`
	// Automatic restart after crash (currently ignored).
	AutoStart int8 `json:"autostart,omitempty"`
	// Amount of target RAM for the VM in MiB. Using zero disables the ballon driver.
	Balloon int `json:"balloon,omitempty"`
	// Select BIOS implementation.
	BIOS string `json:"bios,omitempty"`
	// boot order. ";" separated. : 'order=device1;device2;device3'
	Boot string `json:"boot,omitempty"`
	// Override I/O bandwidth limit (in KiB/s).
	BWLimit int `json:"bwlimit,omitempty"`
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
	// cloud-init: do an automatic package upgrade after the first boot.
	CiUpgrate int8 `json:"ciupgrade,omitempty"`
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
	CpuUnits int `json:"cpuunits,omitempty"`
	// Description for the VM. Shown in the web-interface VM's summary.
	// This is saved as comment inside the configuration file.
	Description string `json:"description,omitempty"`
	// Configure a disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB
	// to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default
	// EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the 'import-from'
	// parameter to import from an existing volume.
	EfiDisk0 int8 `json:"efidisk0,omitempty"`
	// Allow to overwrite existing VM.
	Force int8 `json:"force,omitempty"`
	// Freeze CPU at startup (use 'c' monitor command to start execution).
	Freeze int8 `json:"freeze,omitempty"`
	// Script that will be executed during various steps in the vms lifetime.
	HookScript string `json:"hookscript,omitempty"`
	HostPci
	// Selectively enable hotplug features. This is a comma separated list of hotplug features: 'network', 'disk', 'cpu', 'memory', 'usb' and 'cloudinit'.
	// Use '0' to disable hotplug completely. Using '1' as value is an alias for the default `network,disk,usb`.
	// USB hotplugging is possible for guests with machine version >= 7.1 and ostype l26 or windows > 7.
	HotPlug string `json:"hotplug,omitempty"`
	// Enable/disable hugepages memory.
	HugePages string `json:"hugepages,omitempty"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3).
	// Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	// Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ide
	IPConfig
	// Inter-VM shared memory. Useful for direct communication between VMs, or to the host.
	IVshMem       string `json:"ivshmem,omitempty"`
	KeepHugePages int8   `json:"keephugepages,omitempty"`
	Keyboard      string `json:"keyboard,omitempty"`
	// enable/disable KVM hardware virtualization
	KVM         int8   `json:"kvm,omitempty"`
	LiveRestore int8   `json:"live-restore,omitempty"`
	LocalTime   int8   `json:"localtime,omitempty"`
	Lock        string `json:"lock,omitempty"`
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
	Net
	// node name
	Node string `json:"-"`
	Numa int8   `json:"numa,omitempty"`
	NumaS
	// specifies whether a VM will be started during system bootup
	OnBoot int8 `json:"onboot,omitempty"`
	// quest OS
	OSType OSType `json:"ostype,omitempty"`
	Parallel
	Pool       string `json:"pool,omitempty"`
	Protection int8   `json:"protection,omitempty"`
	// Allow reboot. if set to '0' the VM exit on reboot
	Reboot int    `json:"reboot,omitempty"`
	RNG0   string `json:"rng0,omitempty"`
	Sata
	// use volume as scsi hard disk or cd-rom
	// use special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume
	// use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Scsi
	// SCSI controller model
	ScsiHw ScsiHw `json:"scsihw,omitempty"`
	// cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	SearchDomain string `json:"searchdomain,omitempty"`
	Serial
	Shares  int    `json:"shares,omitempty"`
	SMBios1 string `json:"smbios1,omitempty"`
	SMP     int    `json:"smp,omitempty"`
	// number of sockets
	Sockets           int    `json:"sockets,omitempty"`
	SpiceEnhancements string `json:"spice_enhancements,omitempty"`
	// cloud-init setup public ssh keys (one key per line, OpenSSH format)
	SSHKeys   string `json:"sshkeys,omitempty"`
	StartDate string `json:"startdate,omitempty"`
	StartUp   string `json:"startup,omitempty"`
	Storage   string `json:"storage,omitempty"`
	Tablet    int8   `json:"tablet,omitempty"`
	// tags of the VM. only for meta information
	Tags string `json:"tags,omitempty"`
	TDF  int8   `json:"tdf,omitempty"`
	// enable/disable template
	Template  int8   `json:"template,omitempty"`
	TPMState0 string `json:"tpmstate,omitempty"`
	UnUsed
	USB
	VCPUs int    `json:"vcpus,omitempty"`
	VGA   string `json:"vga,omitempty"`
	VirtIO
	VMGenID        string `json:"vmgenid,omitempty"`
	VMID           *int   `json:"vmid,omitempty"`
	VMStateStorage string `json:"vmstatestorage,omitempty"`
	WatchDog       string `json:"watchdog,omitempty"`
}

type VirtualMachineConfig struct {
	// Enable/disable ACPI.
	ACPI int8 `json:"acpi,omitempty"`
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
	AutoStart int8 `json:"autostart,omitempty"`
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
	CpuUnits    int    `json:"cpuunits,omitempty"`
	Description string `json:"description,omitempty"`
	EfiDisk0    int8   `json:"efidisk0,omitempty"`
	Freeze      int8   `json:"freeze,omitempty"`
	HookScript  string `json:"hookscript,omitempty"`
	HostPci     `json:",inline"`
	HotPlug     string `json:"hotplug,omitempty"`
	HugePages   string `json:"hugepages,omitempty"`
	// Use volume as IDE hard disk or CD-ROM (n is 0 to 3).
	// Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume.
	// Use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Ide           `json:",inline"`
	IPConfig      `json:",inline"`
	IvshMem       string `json:"ivshmem,omitempty"`
	KeepHugePages int8   `json:"keephugepages,omitempty"`
	Keyboard      string `json:"keyboard,omitempty"`
	// enable/disable KVM hardware virtualization
	Kvm       int8   `json:"kvm,omitempty"`
	LocalTime int8   `json:"localtime,omitempty"`
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
	Net   `json:",inline"`
	Numa  int8 `json:"numa,omitempty"`
	NumaS `json:",inline"`
	// specifies whether a VM will be started during system bootup
	OnBoot int8 `json:"onboot,omitempty"`
	// quest OS
	OSType     OSType `json:"ostype,omitempty"`
	Parallel   `json:",inline"`
	Protection int8 `json:"protection,omitempty"`
	// Allow reboot. if set to '0' the VM exit on reboot
	Reboot int    `json:"reboot,omitempty"`
	RNG0   string `json:"rng0,omitempty"`
	Sata   `json:",inline"`
	// use volume as scsi hard disk or cd-rom
	// use special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume
	// use STORAGE_ID:0 and the 'import-from' parameter to import from an existing volume.
	Scsi `json:",inline"`
	// SCSI controller model
	ScsiHw ScsiHw `json:"scsihw,omitempty"`
	// cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set.
	SearchDomain string `json:"searchdomain,omitempty"`
	Serial       `json:",inline"`
	Shares       int    `json:"shares,omitempty"`
	SMBios1      string `json:"smbios1,omitempty"`
	SMP          int    `json:"smp,omitempty"`
	// number of sockets
	Sockets           int    `json:"sockets,omitempty"`
	SpiceEnhancements string `json:"spice_enhancements,omitempty"`
	// cloud-init setup public ssh keys (one key per line, OpenSSH format)
	SSHKeys   string `json:"sshkeys,omitempty"`
	StartDate string `json:"startdate,omitempty"`
	StartUp   int8   `json:"startup,omitempty"`
	Tablet    int8   `json:"tablet,omitempty"`
	// tags of the VM. only for meta information
	Tags string `json:"tags,omitempty"`
	TDF  int8   `json:"tdf,omitempty"`
	// enable/disable template
	Template       int8   `json:"template,omitempty"`
	TPMState0      string `json:"tpmstate,omitempty"`
	UnUsed         `json:",inline"`
	VCPUs          int    `json:"vcpus,omitempty"`
	VGA            string `json:"vga,omitempty"`
	VirtIO         `json:",inline"`
	VMGenID        string `json:"vmgenid,omitempty"`
	VMStateStorage string `json:"vmstatestorage,omitempty"`
	WatchDog       string `json:"watchdog,omitempty"`
}

type VirtualMachineRebootOption struct {
	TimeOut int `json:"timeout,omitempty"`
}

type VirtualMachineResumeOption struct {
	NoCheck  int8 `json:"nocheck,omitempty"`
	SkipLock int8 `json:"skiplock,omitempty"`
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
	SkipLock      int8   `json:"skiplock,omitempty"`
	// some command save/restore state from this location
	StateURI string `json:"stateuri,omitempty"`
	// Mapping from source to target storages. Providing only a single storage ID maps all source storages to that storage.
	// Providing the special value '1' will map each source storage to itself.
	TargetStoraage string `json:"targetstorage,omitempty"`
	TimeOut        int    `json:"timeout,omitempty"`
}

type VirtualMachineStopOption struct {
	KeepActive   int8   `json:"keepActive,omitempty"`
	MigratedFrom string `json:"migratedfrom,omitempty"`
	SkipLock     int8   `json:"skiplock,omitempty"`
	TimeOut      int    `json:"timeout,omitempty"`
}
