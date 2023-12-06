package rest

import (
	"context"
	"fmt"

	"github.com/k8s-proxmox/proxmox-go/api"
)

func (c *RESTClient) GetTasks(ctx context.Context, node string) ([]*api.Task, error) {
	var tasks []*api.Task
	path := fmt.Sprintf("/nodes/%s/tasks", node)
	if err := c.Get(ctx, path, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (c *RESTClient) GetTask(ctx context.Context, node string, upid string) (*api.Task, error) {
	tasks, err := c.GetTasks(ctx, node)
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if task.UPID == upid {
			return task, nil
		}
	}
	return nil, NotFoundErr
}
