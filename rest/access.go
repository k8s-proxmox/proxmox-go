package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) PostTicket(ctx context.Context, req TicketRequest) (*api.Session, error) {
	var session *api.Session
	if err := c.Post(ctx, "/access/ticket", req, &session); err != nil {
		return nil, err
	}
	return session, nil
}

func retrieveSessionTokens(baseUrl string, req TicketRequest) (*api.Session, error) {
	endpoint := baseUrl + "/access/ticket"
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(jsonReq)
	httpReq, err := http.NewRequestWithContext(context.Background(), http.MethodPost, endpoint, body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	buf, err := checkResponse(httpRsp)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session token: %v", err)
	}
	var datakey map[string]json.RawMessage
	if err := json.Unmarshal(buf, &datakey); err != nil {
		return nil, err
	}
	var session *api.Session
	data := datakey["data"]
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}
	return session, nil
}
