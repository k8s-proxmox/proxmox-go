package proxmox

import (
	"context"
	"testing"
	"time"
)

func (s *TestSuite) TestVNCWebSocketClient() {
	testNode := s.getTestNode()
	client, err := s.service.NewNodeVNCWebSocketConnection(context.TODO(), testNode.Node)
	if err != nil {
		s.T().Fatalf("failed to create new vnc client: %v", err)
	}
	defer client.Close()

	if err := client.Write("pwd"); err != nil {
		s.T().Fatalf("write error: %v", err)
	}

	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	out, _, err := client.Read(ctx)
	if err != nil {
		s.T().Fatalf("failed read message: %v", err)
	}

	s.T().Logf("read message: %s", out)
}

func (s *TestSuite) TestExec() {
	testNode := s.getTestNode()
	client, err := s.service.NewNodeVNCWebSocketConnection(context.TODO(), testNode.Node)
	if err != nil {
		s.T().Fatalf("failed to create new vnc client: %v", err)
	}
	defer client.Close()

	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	out, code, err := client.Exec(ctx, "whoami | base64 | base64 -d")
	if err != nil {
		s.T().Fatalf("failed to exec command: %s : %d : %v", out, code, err)
	}
	s.T().Logf("exec command : %s : %d", out, code)
}

func (s *TestSuite) TestWriteFile() {
	testNode := s.getTestNode()
	client, err := s.service.NewNodeVNCWebSocketConnection(context.TODO(), testNode.Node)
	if err != nil {
		s.T().Fatalf("failed to create new vnc client: %v", err)
	}
	defer client.Close()

	ctx, _ := context.WithTimeout(context.TODO(), 15*time.Second)
	err = client.WriteFile(ctx, "this is a file content", "~/test-write-file.txt")
	if err != nil {
		s.T().Fatalf("failed to exec command: %v", err)
	}
}

func TestParseFinMessage(t *testing.T) {
	testMsg := " daf" + finMessage + "123\n"
	if parseFinMessage(testMsg) == "" {
		t.Fatalf("wrong")
	}
}
