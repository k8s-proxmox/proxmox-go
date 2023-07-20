package rest

import (
	"fmt"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetVirtualMachines(node string) ([]*api.VirtualMachine, error) {
	path := fmt.Sprintf("/nodes/%s/qemu", node)
	var vms []*api.VirtualMachine
	if err := c.Get(path, &vms); err != nil {
		return nil, err
	}
	return vms, nil
}

func (c *RESTClient) GetVirtualMachine(node string, vmid int) (*api.VirtualMachine, error) {
	vms, err := c.GetVirtualMachines(node)
	if err != nil {
		return nil, err
	}
	for _, vm := range vms {
		if vm.VMID == vmid {
			return vm, nil
		}
	}
	return nil, NotFoundErr
}

func (c *RESTClient) CreateVirtualMachine(node string, vmid int, options api.VirtualMachineCreateOptions) (*string, error) {
	options.VMID = vmid
	path := fmt.Sprintf("/nodes/%s/qemu", node)
	var upid *string
	if err := c.Post(path, options, &upid); err != nil {
		return nil, err
	}
	return upid, nil
}

func (c *RESTClient) DeleteVirtualMachine(node string, vmid int) (*string, error) {
	path := fmt.Sprintf("/nodes/%s/qemu/%d", node, vmid)
	var upid *string
	if err := c.Delete(path, nil, upid); err != nil {
		return nil, err
	}
	return upid, nil
}

func (c *RESTClient) GetVirtualMachineConfig(node string, vmid int) (*api.VirtualMachineConfig, error) {
	path := fmt.Sprintf("/nodes/%s/qemu/%d/config", node, vmid)
	var config *api.VirtualMachineConfig
	if err := c.Get(path, &config); err != nil {
		return nil, err
	}
	return config, nil
}
