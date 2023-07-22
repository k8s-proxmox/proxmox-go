package rest

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) PostTicket(ctx context.Context, req TicketRequest) (*api.Session, error) {
	if err := c.Post(ctx, "/access/ticket", req, &c.session); err != nil {
		return nil, err
	}
	return c.session, nil
}
