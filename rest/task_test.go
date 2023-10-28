package rest

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) GetTestNode() *api.Node {
	nodes, err := s.restclient.GetNodes(context.TODO())
	if err != nil {
		s.T().Errorf("failed to get nodes: %v", err)
	}
	return nodes[0]
}

func (s *TestSuite) TestGetTasks() {
	testNodeName := s.GetTestNode().Node

	tasks, err := s.restclient.GetTasks(context.TODO(), testNodeName)
	if err != nil {
		s.T().Errorf("failed to get tasks on a node=%s: %v", testNodeName, err)
	}
	s.T().Logf("get tasks: %v", tasks)
}

func (s *TestSuite) TestGetTask() {
	testNodeName := s.GetTestNode().Node

	tasks, err := s.restclient.GetTasks(context.TODO(), testNodeName)
	if err != nil {
		s.T().Errorf("failed to get tasks on a node=%s: %v", testNodeName, err)
	}

	testTaskUPID := tasks[0].UPID
	task, err := s.restclient.GetTask(context.TODO(), testNodeName, testTaskUPID)
	if err != nil {
		s.T().Errorf("failed to get task(upid=%s) on a node=%s: %v", testTaskUPID, testNodeName, err)
	}
	s.T().Logf("get task: %v", *task)
}
