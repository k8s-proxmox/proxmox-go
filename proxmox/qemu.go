package proxmox

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/sp-yduck/proxmox-go/api"
	"github.com/sp-yduck/proxmox-go/rest"
)

type VirtualMachine struct {
	service    *Service
	restclient *rest.RESTClient
	Node       string
	VM         *api.VirtualMachine
	config     *api.VirtualMachineConfig
}

const (
	UUIDFormat = `[a-f\d]{8}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{12}`
)

// VirtualMachines returns all qemus across all proxmox nodes
func (s *Service) VirtualMachines(ctx context.Context) ([]*api.VirtualMachine, error) {
	nodes, err := s.Nodes(ctx)
	if err != nil {
		return nil, err
	}
	var vms []*api.VirtualMachine
	for _, node := range nodes {
		v, err := s.restclient.GetVirtualMachines(ctx, node.Node)
		if err != nil {
			return nil, err
		}
		vms = append(vms, v...)
	}
	return vms, nil
}

func (s *Service) VirtualMachine(ctx context.Context, vmid int) (*VirtualMachine, error) {
	nodes, err := s.Nodes(ctx)
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		vm, err := s.restclient.GetVirtualMachine(ctx, node.Node, vmid)
		if err != nil {
			if rest.IsNotFound(err) {
				continue
			}
			return nil, err
		}
		return &VirtualMachine{service: s, restclient: s.restclient, VM: vm, Node: node.Node}, nil
	}
	return nil, rest.NotFoundErr
}

func (s *Service) CreateVirtualMachine(ctx context.Context, node string, vmid int, option api.VirtualMachineCreateOptions) (*VirtualMachine, error) {
	taskid, err := s.restclient.CreateVirtualMachine(ctx, node, vmid, option)
	if err != nil {
		return nil, err
	}
	if err := s.EnsureTaskDone(ctx, node, *taskid); err != nil {
		return nil, err
	}
	return s.VirtualMachine(ctx, vmid)
}

func (s *Service) VirtualMachineFromUUID(ctx context.Context, uuid string) (*VirtualMachine, error) {
	nodes, err := s.Nodes(ctx)
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		vms, err := s.restclient.GetVirtualMachines(ctx, node.Node)
		if err != nil {
			continue
		}
		for _, vm := range vms {
			config, err := s.restclient.GetVirtualMachineConfig(ctx, node.Node, vm.VMID)
			if err != nil {
				continue
			}
			vmuuid, err := ConvertSMBiosToUUID(config.SMBios1)
			if err != nil {
				continue
			}
			if vmuuid == uuid {
				return &VirtualMachine{service: s, restclient: s.restclient, VM: vm, Node: node.Node, config: config}, nil
			}
		}
	}
	return nil, rest.NotFoundErr
}

func ConvertSMBiosToUUID(smbios string) (string, error) {
	re := regexp.MustCompile(fmt.Sprintf("uuid=%s", UUIDFormat))
	match := re.FindString(smbios)
	if match == "" {
		return "", errors.New("failed to fetch uuid form smbios")
	}
	// match: uuid=<uuid>
	return strings.Split(match, "=")[1], nil
}

func (c *VirtualMachine) Delete(ctx context.Context) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%d", c.Node, c.VM.VMID)
	var upid string
	if err := c.restclient.Delete(ctx, path, nil, &upid); err != nil {
		return err
	}
	return c.service.EnsureTaskDone(ctx, c.Node, upid)
}

func (c *VirtualMachine) GetConfig(ctx context.Context) (*api.VirtualMachineConfig, error) {
	if c.config != nil {
		return c.config, nil
	}
	config, err := c.restclient.GetVirtualMachineConfig(ctx, c.Node, c.VM.VMID)
	if err != nil {
		return nil, err
	}
	c.config = config
	return c.config, err
}

func (c *VirtualMachine) GetOSInfo(ctx context.Context) (*api.OSInfo, error) {
	var osInfo *api.OSInfo
	path := fmt.Sprintf("/nodes/%s/qemu/%d/agent/get-osinfo", c.Node, c.VM.VMID)
	if err := c.restclient.Get(ctx, path, &osInfo); err != nil {
		return nil, err
	}
	return osInfo, nil
}

// size : The new size. With the `+` sign the value is added to the actual size of the volume
// and without it, the value is taken as an absolute one.
// Shrinking disk size is not supported.
// size format : \+?\d+(\.\d+)?[KMGT]?j
func (c *VirtualMachine) ResizeVolume(ctx context.Context, disk, size string) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/resize", c.Node, c.VM.VMID)
	request := make(map[string]interface{})
	request["disk"] = disk
	request["size"] = size
	var v interface{}
	if err := c.restclient.Put(ctx, path, request, &v); err != nil {
		return err
	}
	return nil
}

func (c *VirtualMachine) Start(ctx context.Context, option api.VirtualMachineStartOption) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/status/start", c.Node, c.VM.VMID)
	var upid string
	if err := c.restclient.Post(ctx, path, option, &upid); err != nil {
		return err
	}
	return c.service.EnsureTaskDone(ctx, c.Node, upid)
}

func (c *VirtualMachine) Stop(ctx context.Context, option api.VirtualMachineStopOption) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/status/stop", c.Node, c.VM.VMID)
	var upid string
	if err := c.restclient.Post(ctx, path, option, &upid); err != nil {
		return err
	}
	return c.service.EnsureTaskDone(ctx, c.Node, upid)
}

func (c *VirtualMachine) Resume(ctx context.Context, option api.VirtualMachineResumeOption) error {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/status/resume", c.Node, c.VM.VMID)
	var upid string
	if err := c.restclient.Post(ctx, path, option, &upid); err != nil {
		return err
	}
	return c.service.EnsureTaskDone(ctx, c.Node, upid)
}
