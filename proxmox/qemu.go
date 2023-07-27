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
		return &VirtualMachine{restclient: s.restclient, VM: vm, Node: node.Node}, nil
	}
	return nil, rest.NotFoundErr
}

func (s *Service) VirtualMachineFromUUID(ctx context.Context, uuid string) (*VirtualMachine, error) {
	nodes, err := s.Nodes(ctx)
	if err != nil {
		return nil, err
	}
	for _, node := range nodes {
		vms, err := s.restclient.GetVirtualMachines(ctx, node.Node)
		if err != nil {
			return nil, err
		}
		for _, vm := range vms {
			config, err := s.restclient.GetVirtualMachineConfig(ctx, node.Node, vm.VMID)
			if err != nil {
				return nil, err
			}
			vmuuid, err := convertSMBiosToUUID(config.SMBios1)
			if err != nil {
				return nil, err
			}
			if vmuuid == uuid {
				return &VirtualMachine{restclient: s.restclient, VM: vm, Node: node.Node, config: config}, nil
			}
		}
	}
	return nil, rest.NotFoundErr
}

func convertSMBiosToUUID(smbios string) (string, error) {
	re := regexp.MustCompile(fmt.Sprintf("uuid=%s", UUIDFormat))
	match := re.FindString(smbios)
	if match == "" {
		return "", errors.New("failed to fetch uuid form smbios")
	}
	// match: uuid=<uuid>
	return strings.Split(match, "=")[1], nil
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
