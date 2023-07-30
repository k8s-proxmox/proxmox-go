package rest

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (s *TestSuite) TestDialNodeVNCWebSocket() {
	testNode := s.GetTestNode()
	termProxy, err := s.restclient.CreateNodeTermProxy(context.TODO(), testNode.Node, api.TermProxyOption{})
	if err != nil {
		s.T().Fatalf("failed to create termproxy: %v", err)
	}
	s.T().Logf("create termproxy: %v", termProxy)

	conn, err := s.restclient.DialNodeVNCWebSocket(context.TODO(), testNode.Node, *termProxy)
	if err != nil {
		s.T().Fatalf("failed to connect with vncwebsocket: %v", err)
	}
	defer conn.Close()
}
