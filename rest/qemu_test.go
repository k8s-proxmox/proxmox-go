package rest

import (
	"context"

	"github.com/k8s-proxmox/proxmox-go/api"
)

func (s *TestSuite) TestGetVirtualMachines() {
	nodeName := s.GetTestNode().Node
	vms, err := s.restclient.GetVirtualMachines(context.TODO(), nodeName)
	if err != nil {
		s.T().Fatalf("failed to get vms: %v", err)
	}
	s.T().Logf("get vms: %v", vms)
}

func (s *TestSuite) GetTestVM() *api.VirtualMachine {
	nodeName := s.GetTestNode().Node
	vms, err := s.restclient.GetVirtualMachines(context.TODO(), nodeName)
	if err != nil {
		s.T().Fatalf("failed to get vms: %v", err)
	}
	return vms[0]
}

func (s *TestSuite) TestGetVirtualMachine() {
	nodeName := s.GetTestNode().Node
	vmid := s.GetTestVM().VMID
	vm, err := s.restclient.GetVirtualMachine(context.TODO(), nodeName, vmid)
	if err != nil {
		s.T().Fatalf("failed to get vm: %v", err)
	}
	s.T().Logf("get vm: %v", *vm)
}

func (s *TestSuite) TestGetVirtualMachineConfig() {
	nodeName := s.GetTestNode().Node
	vmid := s.GetTestVM().VMID
	config, err := s.restclient.GetVirtualMachineConfig(context.TODO(), nodeName, vmid)
	if err != nil {
		s.T().Fatalf("failed to get vm: %v", err)
	}
	s.T().Logf("get vm config: %v", *config)
}

func (s *TestSuite) TestSetVirtualMachineConfigAsync() {
	nodeName := "assam"
	vmid := 999
	config := api.VirtualMachineConfig{
		CiPassword: "pve",
		CiUser:     "pve",
	}
	upid, err := s.restclient.SetVirtualMachineConfigAsync(context.TODO(), nodeName, vmid, config)
	if err != nil {
		s.T().Fatalf("failed to set vm: %v", err)
	}
	s.T().Logf("set vm config: %s", *upid)

}

func (s *TestSuite) TestCreateVirtualMachineClone() {
	nodeName := s.GetTestNode().Node
	vmid := s.GetTestVM().VMID
	option := api.VirtualMachineCloneOption{}
	upid, err := s.restclient.CreateVirtualMachineClone(context.TODO(), nodeName, vmid, 999, option)
	if err != nil {
		s.T().Fatalf("failed to clone vm: %v", err)
	}
	s.T().Logf("clone vm: %s", *upid)
}
