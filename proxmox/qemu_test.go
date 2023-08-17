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

func (s *TestSuite) TestCreateVirtualMachine() {
	// testNode := s.getTestNode()
	// testVMID, err := s.service.RESTClient().GetNextID(context.TODO())
	// if err != nil {
	// 	s.T().Fatalf("failed to get next id: %v", err)
	// }
	option := api.VirtualMachineCreateOptions{}
	s.T().Logf("option : %v", option)
	// vm, err := s.service.CreateVirtualMachine(context.TODO(), testNode.Node, testVMID, option)
	// if err != nil {
	// 	s.T().Fatalf("failed to create vm: %v", err)
	// }
	// s.T().Logf("create vm : %v", *vm)
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
	testNode := s.getTestNode()
	vms, err := s.service.restclient.GetVirtualMachines(context.TODO(), testNode.Node)
	if err != nil {
		s.T().Fatalf("failed to get vms: %v", err)
	}
	vm, err := s.service.VirtualMachine(context.TODO(), vms[0].VMID)
	if err != nil {
		s.T().Fatalf("failed to get vm: %v", err)
	}
	return nil, vm
}

func (s *TestSuite) TestGetConfig() {
	_, vm := s.getTestVirtualMachine()
	config, err := vm.GetConfig(context.TODO())
	if err != nil {
		s.T().Fatalf("failed to get vm config: %v", err)
	}
	s.T().Logf("get vm config: %v", *config)
}
