package proxmox

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) TestVirtualMachine() {
	_, testVM := s.getTestVirtualMachine()
	vm, err := s.service.VirtualMachine(context.TODO(), testVM.VM.VMID)
	if err != nil {
		s.T().Fatalf("failed to get vm(vmid=%d): %v", testVM.VM.VMID, err)
	}
	s.Assert().Equal(*vm, *testVM)
}

func (s *TestSuite) TestStop() {
	_, testVM := s.getTestVirtualMachine()
	if err := testVM.Stop(context.TODO(), api.VirtualMachineStopOption{}); err != nil {
		s.T().Fatalf("failed to stop vm: %v", err)
	}
}

func (s *TestSuite) getTestNode() *api.Node {
	nodes, err := s.service.restclient.GetNodes(context.TODO())
	if err != nil {
		s.T().Fatalf("failed to get nodes: %v", err)
	}
	return nodes[0]
}

func (s *TestSuite) getTestVirtualMachine() (*api.Node, *VirtualMachine) {
	node := s.getTestNode()
	vms, err := s.service.restclient.GetVirtualMachines(context.TODO(), node.Node)
	if err != nil {
		s.T().Fatalf("failed to get vms: %v", err)
	}
	vm, err := s.service.VirtualMachine(context.TODO(), vms[0].VMID)
	if err != nil {
		s.T().Fatalf("failed to get vm: %v", err)
	}
	return node, vm
}
