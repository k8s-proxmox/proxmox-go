package rest

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) PostTicket(ctx context.Context, req TicketRequest) (*api.Session, error) {
	var session *api.Session
	if err := c.Post(ctx, "/access/ticket", req, &session); err != nil {
		return nil, err
	}
	return session, nil
}
