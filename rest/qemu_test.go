package rest

import "github.com/sp-yduck/proxmox-go/api"

func (s *TestSuite) TestGetVirtualMachines() {
	nodeName := s.GetTestNode().Node
	vms, err := s.restclient.GetVirtualMachines(nodeName)
	if err != nil {
		s.T().Fatalf("failed to get vms: %v", err)
	}
	s.T().Logf("get vms: %v", vms)
}

func (s *TestSuite) GetTestVM() *api.VirtualMachine {
	nodeName := s.GetTestNode().Node
	vms, err := s.restclient.GetVirtualMachines(nodeName)
	if err != nil {
		s.T().Fatalf("failed to get vms: %v", err)
	}
	return vms[0]
}

func (s *TestSuite) TestGetVirtualMachine() {
	nodeName := s.GetTestNode().Node
	vmid := s.GetTestVM().VMID
	vm, err := s.restclient.GetVirtualMachine(nodeName, vmid)
	if err != nil {
		s.T().Fatalf("failed to get vm: %v", err)
	}
	s.T().Logf("get vm: %v", *vm)
}

func (s *TestSuite) TestGetVirtualMachineConfig() {
	nodeName := s.GetTestNode().Node
	vmid := s.GetTestVM().VMID
	config, err := s.restclient.GetVirtualMachineConfig(nodeName, vmid)
	if err != nil {
		s.T().Fatalf("failed to get vm: %v", err)
	}
	s.T().Logf("get vm config: %v", *config)
}
