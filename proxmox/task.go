package proxmox

import "fmt"

type Task struct {
	Endtime int `json:"endtime"`
	// Id        string `json:"id,omitempty"`
	Node      string `json:"node"`
	PID       int    `json:"pid"`
	PStart    int    `json:"pstart"`
	StartTime int    `json:"starttime"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UPID      string `json:"upid"`
	User      string `json:"user"`
}

type TaskStatus struct {
	Exitstatus string `json:"exitstatus"`
	Id         string `json:"id"`
	Node
}

func (c *RESTClient) GetTasks(node string) ([]*Task, error) {
	var tasks []*Task
	path := fmt.Sprintf("/nodes/%s/tasks", node)
	if err := c.Get(path, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (c *RESTClient) GetTask(node string, upid string) (*Task, error) {
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
