package rest

import (
	"context"
	"encoding/json"
)

func (c *RESTClient) GetNextID(ctx context.Context) (int, error) {
	var res json.Number
	if err := c.Get(ctx, "/cluster/nextid", &res); err != nil {
		return 0, err
	}
	nextid, err := res.Int64()
	if err != nil {
		return 0, err
	}
	return int(nextid), nil
}
