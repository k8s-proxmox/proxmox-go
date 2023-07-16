package proxmox

import (
	"fmt"
)

func (c *RESTClient) GetVirtualMachines(node string) ([]*VirtualMachine, error) {
	path := fmt.Sprintf("/nodes/%s/qemu", node)
	var vms []*VirtualMachine
	if err := c.Get(path, &vms); err != nil {
		return nil, err
	}
	return vms, nil
}

func (c *RESTClient) GetVirtualMachine(node string, vmid int) (*VirtualMachine, error) {
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

func (c *RESTClient) CreateVirtualMachine(node string, vmid int, options VirtualMachineCreateOptions) (*string, error) {
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
