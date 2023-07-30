package proxmox

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/gorilla/websocket"
	"github.com/sp-yduck/proxmox-go/api"
)

const (
	finMessage       = "done with status: "
	finMessageFormat = finMessage + `[0-9]+`
)

type VNCWebSocketClient struct {
	conn *websocket.Conn
}

func (s *Service) NewNodeVNCWebSocketConnection(ctx context.Context, nodeName string) (*VNCWebSocketClient, error) {
	termProxy, err := s.restclient.CreateNodeTermProxy(ctx, nodeName, api.TermProxyOption{})
	if err != nil {
		return nil, err
	}
	conn, err := s.restclient.DialNodeVNCWebSocket(ctx, nodeName, *termProxy)
	if err != nil {
		return nil, err
	}

	return &VNCWebSocketClient{conn: conn}, nil
}

func (c *VNCWebSocketClient) Close() {
	c.conn.Close()
}

func (c *VNCWebSocketClient) Write(cmd string) error {
	b := []byte(fmt.Sprintf("%s\n", cmd))
	bheader := []byte(fmt.Sprintf("0:%d:", len(b)))
	bmsg := append(bheader, b...)
	if err := c.conn.WriteMessage(websocket.BinaryMessage, bmsg); err != nil {
		return err
	}
	return c.sendFinMessage()
}

func (c *VNCWebSocketClient) sendFinMessage() error {
	b := []byte(fmt.Sprintf(`echo "%s$?"%s`, finMessage, "\n"))
	bheader := []byte(fmt.Sprintf("0:%d:", len(b)))
	bmsg := append(bheader, b...)
	if err := c.conn.WriteMessage(websocket.BinaryMessage, bmsg); err != nil {
		return err
	}
	return nil
}

// Read() reads message until find fin message
// then returns whole message and status code
func (c *VNCWebSocketClient) Read(ctx context.Context) (outputs string, code int, err error) {
	done := make(chan error, 1)
	go func() {
		defer close(done)
		for {
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				done <- err
				return
			}
			outputs += string(msg)
			finMsg := parseFinMessage(string(msg))
			if finMsg != "" {
				code, err = parseStatusFromFinMessage(finMsg)
				done <- err
				return
			}
		}
	}()
	select {
	case err = <-done:
		return outputs, code, err
	case <-ctx.Done():
		return outputs, -1, errors.New("context deadline exceeded")
	}
}

// Exec executes a command and return error if code is not 0
// usually out contains many extra messages that is just useless
func (c *VNCWebSocketClient) Exec(ctx context.Context, cmd string) (out string, code int, err error) {
	if err := c.Write(cmd); err != nil {
		return "", 0, err
	}
	out, code, err = c.Read(ctx)
	if err != nil {
		return out, code, err
	}
	if code != 0 {
		return out, code, errors.Errorf("exit with non zero code: %d", code)
	}
	return out, 0, nil
}

func parseFinMessage(message string) string {
	re := regexp.MustCompile(finMessageFormat)
	return re.FindString(message)
}

func parseStatusFromFinMessage(message string) (int, error) {
	re := regexp.MustCompile(finMessageFormat)
	match := re.FindString(message)
	if match == "" {
		return 0, errors.Errorf("failed to find status code from %s", message)
	}
	statusCode := strings.Split(match, ": ")[1]
	return strconv.Atoi(statusCode)
}
