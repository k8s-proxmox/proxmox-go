package proxmox

import (
	"context"
	"errors"
	"time"

	"github.com/sp-yduck/proxmox-go/api"
	"github.com/sp-yduck/proxmox-go/rest"
)

const (
	TaskStatusOK = "OK"
)

func (s *Service) MustGetTask(ctx context.Context, node string, upid string) (*api.Task, error) {
	time.Sleep(time.Second * 60)
	for i := 0; i < 40; i++ {
		task, err := s.restclient.GetTask(ctx, node, upid)
		if err != nil {
			if rest.IsNotFound(err) {
				time.Sleep(time.Second * 1)
				continue
			}
			return nil, err
		}
		return task, nil
	}
	return nil, errors.New("task wait deadline exceeded")
}

func (s *Service) EnsureTaskDone(ctx context.Context, node, upid string) error {
	task, err := s.MustGetTask(ctx, node, upid)
	if err != nil {
		return err
	}
	if task.Status != TaskStatusOK {
		return errors.New(task.Status)
	}
	return nil
}
