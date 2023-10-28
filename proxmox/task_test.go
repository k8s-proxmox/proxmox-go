package proxmox

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) TestMustGetTask() {
	testNode, testTask := s.getTestTask()
	task, err := s.service.MustGetTask(context.TODO(), testNode.Node, testTask.UPID)
	if err != nil {
		s.T().Fatalf("failed to get task: %v", err)
	}
	s.Assert().Equal(*task, *testTask)
}

func (s *TestSuite) TestEnsureTaskDone() {
	testNode, testTask := s.getTestTask()
	err := s.service.EnsureTaskDone(context.TODO(), testNode.Node, testTask.UPID)
	if err != nil {
		s.T().Fatalf("failed to get task: %v", err)
	}
}

func (s *TestSuite) getTestTask() (*api.Node, *api.Task) {
	node := s.getTestNode()
	tasks, err := s.service.restclient.GetTasks(context.TODO(), node.Node)
	if err != nil {
		s.T().Fatalf("failed to get tasks: %v", err)
	}
	return node, tasks[0]
}
