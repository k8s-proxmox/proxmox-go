package proxmox

func (c *RESTClient) PostTicket(req TicketRequest) (*Session, error) {
	if err := c.Post("/access/ticket", req, &c.session); err != nil {
		return nil, err
	}
	return c.session, nil
}
