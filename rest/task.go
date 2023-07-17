package rest

import (
	"fmt"

	"github.com/sp-yduck/proxmox-go/api"
)

func (c *RESTClient) GetTasks(node string) ([]*api.Task, error) {
	var tasks []*api.Task
	path := fmt.Sprintf("/nodes/%s/tasks", node)
	if err := c.Get(path, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (c *RESTClient) GetTask(node string, upid string) (*api.Task, error) {
	tasks, err := c.GetTasks(node)
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
