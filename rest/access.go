package rest

import (
	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) PostTicket(req TicketRequest) (*api.Session, error) {
	if err := c.Post("/access/ticket", req, &c.session); err != nil {
		return nil, err
	}
	return c.session, nil
}
