package proxmox

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gorilla/websocket"
	"github.com/sp-yduck/proxmox-go/api"
)

const (
	finMessage       = "done with status: "
	finMessageFormat = finMessage + `[0-9]+`
)

type VNCWebSocketClient struct {
	conn   *websocket.Conn
	ticker *time.Ticker
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

	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				conn.WriteMessage(websocket.BinaryMessage, []byte("2"))
			}
		}
	}()

	client := &VNCWebSocketClient{conn: conn, ticker: ticker}

	credentials := s.restclient.Credentials()
	// login is required for non-root user: https://git.proxmox.com/?p=pve-manager.git;a=blob;f=PVE/API2/Nodes.pm;h=0843c3a3c6cee7c763bf4ac9d8b75ab298f1373e;hb=HEAD#l913
	if credentials.Username != "root@pam" {
		err := client.Login(strings.Replace(credentials.Username, "@pam", "", 1), credentials.Password)
		if err != nil {
			return nil, fmt.Errorf("login failed: %v", err)
		}
	}

	return client, nil
}

func (c *VNCWebSocketClient) Close() {
	c.conn.Close()
	c.ticker.Stop()
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

func (c *VNCWebSocketClient) ExpectMessage(expected string) error {
	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("failed to read message: %v", err)
	}

	if !strings.Contains(string(msg), expected) {
		return fmt.Errorf("read expected '%s' did not contain expected '%s'", string(msg), expected)
	}
	return nil
}

func (c *VNCWebSocketClient) ExpectMessageMultiline(expected string) error {
	dataStream := make(chan string, 1)
	readErr := make(chan error)
	stop := make(chan bool)
	go func() {
	loop:
		for {
			select {
			case <-stop:
				break loop
			default:
				var r io.Reader
				_, r, err := c.conn.NextReader()
				if err != nil {
					readErr <- err
					break loop
				}

				br := bufio.NewReader(r)
				s, e := ReadLine(br)
				for e == nil {
					dataStream <- s
					s, e = ReadLine(br)
				}
				if e == io.EOF {
					continue loop
				} else {
					readErr <- e
				}
				break loop

			}
		}
		close(dataStream)
		close(readErr)
	}()
	defer close(stop)

	for {
		select {
		case data := <-dataStream:
			if strings.Contains(data, expected) {
				stop <- true
				return nil
			}
		case <-time.After(time.Duration(10) * time.Second):
			return fmt.Errorf("timeout while reading '%s' from multiline response", expected)
		case err := <-readErr:
			return err
		}
	}
}

func ReadLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func (c *VNCWebSocketClient) WriteMessage(message string) error {
	b := []byte(message)
	bheader := []byte(fmt.Sprintf("0:%d:", len(b)))
	bmsg := append(bheader, b...)
	if err := c.conn.WriteMessage(websocket.BinaryMessage, bmsg); err != nil {
		return fmt.Errorf("failed writing message '%s': %v", message, err)
	}
	return nil
}

func (c *VNCWebSocketClient) Login(username string, password string) error {
	// proxmox task result
	if err := c.ExpectMessage("OK"); err != nil {
		return err
	}
	// linux login prompt
	if err := c.ExpectMessage("login"); err != nil {
		return err
	}

	if err := c.WriteMessage(fmt.Sprintf("%s\n", username)); err != nil {
		return err
	}

	// password prompt for user
	if err := c.ExpectMessage(username); err != nil {
		return err
	}
	if err := c.ExpectMessage("Password"); err != nil {
		return err
	}
	if err := c.WriteMessage(fmt.Sprintf("%s\n", password)); err != nil {
		return err
	}

	// successful login
	if err := c.ExpectMessageMultiline("Last login"); err != nil {
		return err
	}
	return nil
}

func (c *VNCWebSocketClient) WriteFile(ctx context.Context, content, path string) error {
	c.Exec(ctx, fmt.Sprintf("rm %s", path))
	if out, _, err := c.Exec(ctx, fmt.Sprintf("touch %s", path)); err != nil {
		return errors.Wrap(err, out)
	}
	chunks := chunkString(content, 2000)
	for _, chunk := range chunks {
		b64chunk := base64.StdEncoding.EncodeToString([]byte(chunk))
		out, _, err := c.Exec(ctx, fmt.Sprintf("echo %s | base64 -d >> %s", b64chunk, path))
		if err != nil {
			return errors.Wrap(err, out)
		}
	}
	return nil
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)
	if len(runes) == 0 {
		return []string{s}
	}
	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
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
