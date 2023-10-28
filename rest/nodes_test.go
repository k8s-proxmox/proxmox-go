package rest

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) TestCreateTermProxy() {
	testNode := s.GetTestNode()
	termProxy, err := s.restclient.CreateNodeTermProxy(context.TODO(), testNode.Node, api.TermProxyOption{})
	if err != nil {
		s.T().Fatalf("failed to create termproxy: %v", err)
	}
	s.T().Logf("create termproxy: %v", termProxy)
}

func (s *TestSuite) TestGetVNCWebSocket() {
	testNode := s.GetTestNode()
	termProxy, err := s.restclient.CreateNodeTermProxy(context.TODO(), testNode.Node, api.TermProxyOption{})
	if err != nil {
		s.T().Fatalf("failed to create termproxy: %v", err)
	}
	s.T().Logf("create termproxy: %v", termProxy)

	websocket, err := s.restclient.GetNodeVNCWebSocket(context.TODO(), testNode.Node, termProxy.Port, termProxy.Ticket)
	if err != nil {
		s.T().Fatalf("failed to get vncwebsocket: %v", err)
	}
	s.T().Logf("get vncwebsocket: %v", websocket)
}
